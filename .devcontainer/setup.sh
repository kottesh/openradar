#!/usr/bin/env bash
set -euo pipefail

cd "$(git rev-parse --show-toplevel)"

echo "[openradar] Building the application image..."
docker build -t openradar:test .
docker network inspect openradar-net >/dev/null 2>&1 || docker network create openradar-net >/dev/null
docker volume inspect openradar-db-data >/dev/null 2>&1 || docker volume create openradar-db-data >/dev/null

echo "[openradar] Setup complete. Containers start automatically when GITHUB_TOKEN is available."
