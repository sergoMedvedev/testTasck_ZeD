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
ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME

ENV DB_HOST=$DB_HOST \
    DB_PORT=$DB_PORT \
    DB_USER=$DB_USER \
    DB_PASSWORD=$DB_PASSWORD \
    DB_NAME=$DB_NAME

EXPOSE 5000

# Запускаем приложение
CMD ["src/main"]