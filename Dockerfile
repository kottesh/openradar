FROM node:22.20.0-alpine AS frontend-builder

WORKDIR /app/frontend
COPY app/package*.json ./
RUN npm ci
COPY app/ ./
RUN npm run build

FROM golang:1.25.0-alpine AS go-builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=frontend-builder /app/frontend/dist ./app/dist
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o server cmd/server/main.go

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates git \
    && addgroup -S -g 10001 openradar \
    && adduser -S -D -H -u 10001 -G openradar openradar
COPY --from=go-builder --chown=openradar:openradar /app/server ./server
USER openradar
ENV PORT=8080
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:${PORT}/ || exit 1
CMD ["./server"]
