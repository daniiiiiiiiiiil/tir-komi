FROM golang:1.26.4-alpine

WORKDIR /app

# Копируем go.mod из папки backend
COPY backend/go.mod ./
RUN go mod download

# Копируем весь код backend
COPY backend/ .

# Собираем из cmd/app (относительно backend)
RUN go build -o main ./cmd/app

EXPOSE 8080

CMD ["./main"]