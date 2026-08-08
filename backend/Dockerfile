FROM golang:1.26.4-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum для скачивания зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 go build -o /app/exe /app/cmd/app

# Финальный образ
FROM alpine:3.23

WORKDIR /app

# Копируем собранный бинарник
COPY --from=builder /app/exe /app/

# Копируем статические файлы (HTML, CSS, JS, изображения)
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/public ./public
COPY --from=builder /app/*.html ./
# Если есть другие статические директории, добавьте их:
# COPY --from=builder /app/uploads ./uploads

# Открываем порт (замените на ваш порт, если отличается)
EXPOSE 8080

# Запускаем приложение
CMD ["/app/exe"]