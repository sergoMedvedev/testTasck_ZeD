# Используем официальный образ Go для сборки приложения
FROM golang:1.18 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum, чтобы скачать зависимости
COPY ./document/go.mod ./


# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы приложения в контейнер
COPY . .

# Собираем Go приложение
RUN go build -o main .

# Используем минимальный образ для запуска приложения
FROM alpine:latest

# Устанавливаем необходимые зависимости
RUN apk --no-cache add ca-certificates

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /root/

# Копируем собранное приложение из контейнера builder
COPY --from=builder /app/main .

# Устанавливаем переменные окружения для подключения к базе данных
ENV DB_HOST=localhost \
    DB_PORT='8080' \
    DB_USER=postgres \
    DB_PASSWORD=password \
    DB_NAME=test

# Запускаем приложение
CMD ["./main"]