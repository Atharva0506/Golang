# Stage 1 — Builder: download dependencies and compile the API binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Download modules first so this layer is cached when only source changes
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source and build a statically-linked binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /app/api ./cmd/api

# Stage 2 — Runner: minimal scratch image; certs are copied from the builder
# (golang:1.24-alpine ships ca-certificates so no extra apk step is needed)
FROM scratch AS runner

# TLS root certificates for outbound HTTPS calls
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app
COPY --from=builder /app/api ./api

# REST / WebSocket
EXPOSE 8080
# gRPC
EXPOSE 50051

ENTRYPOINT ["./api"]
