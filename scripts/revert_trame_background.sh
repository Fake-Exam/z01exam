#!/usr/bin/env bash
set -Eeuo pipefail

FILE="${1:-style.css}"

if [[ ! -f "$FILE" ]]; then
  echo "[error] Fichier introuvable: $FILE" >&2
  exit 1
fi

if ! grep -q "z02intra-trame-start" "$FILE"; then
  echo "[info] Aucun patch trame detecte dans $FILE"
  exit 0
fi

tmp="$(mktemp)"
awk '
  /\/\* z02intra-trame-start \*\// { skip=1; next }
  /\/\* z02intra-trame-end \*\// { skip=0; next }
  skip==0 { print }
' "$FILE" > "$tmp"

mv "$tmp" "$FILE"
echo "[ok] Patch trame retire de $FILE"
