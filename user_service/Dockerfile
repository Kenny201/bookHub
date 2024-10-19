FROM golang:alpine AS builder

COPY . /app/
WORKDIR /app/

RUN set -Eeux && \
    go mod download && \
    go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o build/user_service ./cmd/user/main.go

FROM scratch
WORKDIR /app

COPY --from=builder /app/build/user_service /app/user_service
COPY --from=builder /app/.env /app/.env

ENTRYPOINT ["/app/user_service"]