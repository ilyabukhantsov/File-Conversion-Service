
FROM golang:1.26 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server

FROM debian:bookworm-slim

ENV DEBIAN_FRONTEND=noninteractive


RUN apt-get update && apt-get install -y --no-install-recommends \
        libreoffice-writer \
        libreoffice-calc \
        libreoffice-impress \
        fonts-dejavu \
        fonts-liberation \
        ca-certificates \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -m -u 1000 app && mkdir -p /data && chown app:app /data

ENV HOME=/home/app \
    DATA_DIR=/data \
    PORT=8080 \
    GIN_MODE=release

COPY --from=builder /server /server

USER app
VOLUME ["/data"]
EXPOSE 8080

CMD ["/server"]