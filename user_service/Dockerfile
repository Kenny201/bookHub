FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

FROM golang:alpine AS dev
WORKDIR /app

COPY --from=builder /app ./

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]
