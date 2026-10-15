#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

PORT="${PORT:-8080}"
USE_DOCKER_RUNNER="${USE_DOCKER_RUNNER:-0}"
OFFICIAL_GO_RUNNER_IMAGE="ghcr.io/01-edu/test-go@sha256:9c57c88e68f8937617a26d805396c775d5a88f43e214b91386986f89a9df13bf"
RUNNER_IMAGE="${RUNNER_GO_DOCKER_IMAGE:-$OFFICIAL_GO_RUNNER_IMAGE}"

export PORT

if [[ "$USE_DOCKER_RUNNER" == "1" ]]; then
  if ! command -v docker >/dev/null 2>&1; then
    echo "Docker introuvable. Lance sans docker ou installe Docker."
    exit 1
  fi
  if ! docker image inspect "$RUNNER_IMAGE" >/dev/null 2>&1; then
    echo "Téléchargement de l'image officielle '$RUNNER_IMAGE'..."
    docker pull "$RUNNER_IMAGE"
  fi
  export RUNNER_MODE=docker
  export RUNNER_GO_DOCKER_IMAGE="$RUNNER_IMAGE"
else
  echo "Attention: RUNNER_MODE=docker est requis pour exécuter les tests."
  echo "Active-le avec: USE_DOCKER_RUNNER=1 ./scripts/dev.sh"
fi

cd "$ROOT_DIR/backend"
echo "Serveur sur http://localhost:$PORT"
exec go run ./cmd/monolith
