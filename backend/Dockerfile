FROM golang:1.26.4-alpine AS builder

WORKDIR /app

# Копируем go.mod
COPY backend/go.mod ./
RUN go mod download

# Копируем весь код
COPY backend/ .

# Собираем приложение
RUN go build -o main ./cmd/app

# Финальный этап
FROM alpine:latest

WORKDIR /app

# Копируем бинарник
COPY --from=builder /app/main .

# Копируем папку assets (ВАЖНО!)
COPY --from=builder /app/public/assets ./public/assets

EXPOSE 8080

CMD ["./main"]