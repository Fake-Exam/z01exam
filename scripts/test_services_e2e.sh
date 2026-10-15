#!/usr/bin/env bash
set -euo pipefail

# Proves the deployed topology, not just individual handlers:
# Catalog -> Runner snapshot, Web proxy -> both public APIs, then Docker Go run.
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
TMP_DIR="$(mktemp -d "${TMPDIR:-/tmp}/z02-services-e2e.XXXXXX")"
PORT_BASE="${E2E_PORT_BASE:-18081}"
CATALOG_PORT="$PORT_BASE"
RUNNER_PORT="$((PORT_BASE + 1))"
WEB_PORT="$((PORT_BASE + 2))"
CATALOG_PID=""
RUNNER_PID=""
WEB_PID=""

cleanup() {
  local status=$?
  [[ -z "$WEB_PID" ]] || kill "$WEB_PID" 2>/dev/null || true
  [[ -z "$RUNNER_PID" ]] || kill "$RUNNER_PID" 2>/dev/null || true
  [[ -z "$CATALOG_PID" ]] || kill "$CATALOG_PID" 2>/dev/null || true
  wait 2>/dev/null || true
  if [[ "$status" -eq 0 ]]; then
    rm -rf "$TMP_DIR"
  else
    printf 'E2E failed; diagnostic files retained in %s\n' "$TMP_DIR" >&2
  fi
}
trap cleanup EXIT

# Refuse occupied ports: an unrelated service must never satisfy our checks.
python3 - "$CATALOG_PORT" "$RUNNER_PORT" "$WEB_PORT" <<'PY'
import socket, sys
listeners = []
for port in sys.argv[1:]:
    listener = socket.socket()
    try:
        listener.bind(("127.0.0.1", int(port)))
    except OSError as error:
        raise SystemExit(f"E2E port {port} unavailable: {error}")
    listeners.append(listener)
PY

wait_for_url() {
  local url="$1"
  local destination="$2"
  local pid="$3"
  for _ in $(seq 1 40); do
    kill -0 "$pid" 2>/dev/null || return 1
    if curl -fsS --connect-timeout 1 --max-time 2 "$url" >"$destination"; then
      kill -0 "$pid" 2>/dev/null
      return
    fi
    sleep 1
  done
  return 1
}

cd "$ROOT_DIR/backend"
go build -o "$TMP_DIR/catalog" ./cmd/catalog
go build -o "$TMP_DIR/runner" ./cmd/runner
go build -o "$TMP_DIR/web" ./cmd/web

LISTEN_ADDR="127.0.0.1:$CATALOG_PORT" "$TMP_DIR/catalog" >"$TMP_DIR/catalog.log" 2>&1 &
CATALOG_PID=$!
wait_for_url "http://127.0.0.1:$CATALOG_PORT/api/exercises?track=go" "$TMP_DIR/catalog.json" "$CATALOG_PID"
grep -q '"id"' "$TMP_DIR/catalog.json"

LISTEN_ADDR="127.0.0.1:$RUNNER_PORT" CATALOG_SERVICE_URL="http://127.0.0.1:$CATALOG_PORT" RUNNER_MODE=docker \
  "$TMP_DIR/runner" >"$TMP_DIR/runner.log" 2>&1 &
RUNNER_PID=$!
for _ in $(seq 1 20); do
  kill -0 "$RUNNER_PID" 2>/dev/null || exit 1
  if curl -fsS -X POST -H 'Content-Type: application/json' \
    -d '{"exerciseId":"unknown","code":"x"}' \
    "http://127.0.0.1:$RUNNER_PORT/api/run" >"$TMP_DIR/runner.json"; then
    break
  fi
  sleep 1
done
grep -q 'Exercice inconnu' "$TMP_DIR/runner.json"

LISTEN_ADDR="127.0.0.1:$WEB_PORT" CATALOG_SERVICE_URL="http://127.0.0.1:$CATALOG_PORT" \
  RUNNER_SERVICE_URL="http://127.0.0.1:$RUNNER_PORT" "$TMP_DIR/web" >"$TMP_DIR/web.log" 2>&1 &
WEB_PID=$!
wait_for_url "http://127.0.0.1:$WEB_PORT/" "$TMP_DIR/web-index.html" "$WEB_PID"
grep -q '<title>' "$TMP_DIR/web-index.html"
wait_for_url "http://127.0.0.1:$WEB_PORT/api/exercises?track=go" "$TMP_DIR/web-catalog.json" "$WEB_PID"
grep -q '"id"' "$TMP_DIR/web-catalog.json"

curl -fsS -X POST -H 'Content-Type: application/json' \
  -d '{"exerciseId":"onlya","code":"package main\n\nimport \"fmt\"\n\nfunc main() { fmt.Print(\"a\") }\n"}' \
  "http://127.0.0.1:$WEB_PORT/api/run" >"$TMP_DIR/web-runner.json"
grep -q '"success":true' "$TMP_DIR/web-runner.json"

printf 'E2E services: OK\n'
