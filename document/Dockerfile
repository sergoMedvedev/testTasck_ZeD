# Используем официальный образ Go для сборки приложения
FROM golang:1.18 AS builder

ENV GOPATH=/temp/gopath
# Устанавливаем рабочую директорию внутри контейнера
WORKDIR .

# Копируем go.mod и go.sum, чтобы скачать зависимости
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы приложения в контейнер
COPY . .

# Собираем Go приложение
RUN cd src/ && go build -o main .

# Устанавливаем переменные окружения для подключения к базе данных
ENV DB_HOST=localhost \
    DB_PORT='8080' \
    DB_USER=postgres \
    DB_PASSWORD=password \
    DB_NAME=test

# Запускаем приложение
CMD ["src/main"]