#!/usr/bin/env bash
set -euo pipefail

cd "$(git rev-parse --show-toplevel)"

if [[ -z "${OPENRADAR_GITHUB_TOKEN:-}" ]]; then
  echo "[openradar] OPENRADAR_GITHUB_TOKEN is unavailable. Add it as a Codespaces secret, then rebuild/restart."
  exit 0
fi

origin="https://${CODESPACE_NAME}-8080.${GITHUB_CODESPACES_PORT_FORWARDING_DOMAIN}"
db_password_file="$HOME/.openradar-db-password"
if [[ ! -f "$db_password_file" ]]; then
  umask 077
  openssl rand -hex 24 > "$db_password_file"
fi
db_password="$(<"$db_password_file")"

docker network inspect openradar-net >/dev/null 2>&1 || docker network create openradar-net >/dev/null
docker volume inspect openradar-db-data >/dev/null 2>&1 || docker volume create openradar-db-data >/dev/null

docker rm -f openradar openradar-db >/dev/null 2>&1 || true

docker run -d \
  --name openradar-db \
  --network openradar-net \
  --restart unless-stopped \
  --security-opt no-new-privileges:true \
  -e POSTGRES_DB=secretscan \
  -e POSTGRES_USER=scanner \
  -e POSTGRES_PASSWORD="$db_password" \
  -v openradar-db-data:/var/lib/postgresql/data \
  postgres:17-alpine >/dev/null

for _ in $(seq 1 30); do
  docker exec openradar-db pg_isready -U scanner -d secretscan >/dev/null 2>&1 && break
  sleep 2
done

docker run -d \
  --name openradar \
  --network openradar-net \
  --restart unless-stopped \
  --cap-drop ALL \
  --security-opt no-new-privileges:true \
  --memory 2g \
  --cpus 1 \
  --pids-limit 256 \
  --tmpfs /tmp:rw,nosuid,nodev,noexec,size=1g,uid=10001,gid=10001 \
  -p 127.0.0.1:8080:8080 \
  -e ENV=development \
  -e DATABASE_URL="postgres://scanner:${db_password}@openradar-db:5432/secretscan?sslmode=disable" \
  -e ALLOWED_WEBSOCKET_ORIGINS="$origin,http://localhost:8080,https://localhost:8080" \
  -e SCAN_MAX_REPO_MB=20 \
  -e SCAN_MAX_FILE_KB=256 \
  -e SCAN_MAX_CONCURRENT=2 \
  -e GITHUB_TOKEN="$OPENRADAR_GITHUB_TOKEN" \
  -e PORT=8080 \
  -e WEBHOOK_URL= \
  openradar:test >/dev/null

for _ in $(seq 1 30); do
  [[ "$(docker inspect -f '{{.State.Health.Status}}' openradar 2>/dev/null || true)" == "healthy" ]] && break
  sleep 2
done

echo "[openradar] Running at $origin"
docker ps --filter name=openradar --format 'table {{.Names}}\t{{.Status}}\t{{.Ports}}'
