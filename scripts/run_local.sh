#!/usr/bin/env bash
set -euo pipefail

# Start local postgres via docker-compose for quick dev
if [ "$1" = "up" ]; then
  docker-compose up -d db
  echo "Postgres started. Run make run to start the app."
  exit 0
fi

if [ "$1" = "down" ]; then
  docker-compose down
  exit 0
fi

echo "Usage: $0 {up|down}"
