FROM golang:1.21-alpine

WORKDIR /app

# Копируем только go.mod
COPY go.mod ./
RUN go mod download

# Копируем весь код
COPY . .

# Собираем из cmd/app
RUN go build -o main ./cmd/app

EXPOSE 8080

CMD ["./main"]