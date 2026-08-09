FROM golang:1.26.4-alpine AS builder

WORKDIR /app

# Копируем go.mod
COPY backend/go.mod ./
RUN go mod download

# Копируем весь код
COPY backend/ .

# Собираем приложение
RUN go build -o main ./cmd/app

# Финальный этап - используем тот же alpine, что и в builder
FROM alpine:latest

WORKDIR /app

# Устанавливаем необходимые зависимости (важно!)
RUN apk add --no-cache \
    ca-certificates \
    tzdata

# Копируем бинарник
COPY --from=builder /app/main .

# Копируем папку assets
COPY --from=builder /app/public ./public

# Копируем папку с шаблонами (если она есть)
COPY --from=builder /app/templates ./templates

EXPOSE 8080

CMD ["./main"]