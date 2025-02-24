FROM golang:1.23-alpine AS builder

# Устанавливаем необходимые инструменты для сборки
RUN apk add --no-cache build-base

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код проекта
COPY . .

# Сборка бинарного файла с включенным CGO
RUN CGO_ENABLED=0 go build -o main cmd/api/main.go

# Создаем минимальный образ
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем бинарный файл из builder-образа
COPY --from=builder /app/main .

EXPOSE 1323

# Указываем команду для запуска
CMD ["./main"]
