FROM golang:1.26.4-alpine AS builder

WORKDIR /app

COPY backend/go.mod ./
RUN go mod download

# Копируем ВСЕ файлы из backend
COPY backend/ .

RUN go build -o main ./cmd/app

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

# Копируем бинарник
COPY --from=builder /app/main .

# Копируем папку public со ВСЕМИ файлами
COPY --from=builder /app/public ./public

# Копируем uploads если есть
COPY --from=builder /app/uploads ./uploads

# Проверяем структуру (для отладки)
RUN ls -la && ls -la public/ && ls -la public/assets/ || echo "No assets folder"

EXPOSE 8080

CMD ["./main"]