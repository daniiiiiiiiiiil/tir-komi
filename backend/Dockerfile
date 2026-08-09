FROM golang:1.26.4-alpine

WORKDIR /app

# Копируем go.mod из папки backend
COPY backend/go.mod ./
RUN go mod download

# Копируем код backend
COPY backend/ ./backend/

# Устанавливаем переменную окружения
ENV PROJECT_ROOT=/app/backend

# Собираем приложение
RUN go build -o main ./backend/cmd/app

EXPOSE 8080

CMD ["./main"]