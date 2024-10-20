# Этап сборки
FROM golang:alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем все файлы
COPY . .

# Этап разработки (для live reload)
FROM golang:alpine AS dev
WORKDIR /app

# Копируем файлы из builder
COPY --from=builder /app ./

# Устанавливаем air
RUN go install github.com/air-verse/air@latest

# Запускаем приложение через air с указанием конфигурационного файла
CMD ["air", "-c", ".air.toml"]
