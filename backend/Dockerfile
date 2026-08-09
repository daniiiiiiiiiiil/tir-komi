FROM golang:1.26.4-alpine

WORKDIR /app

# Копируем код и статику в разные места
COPY backend/ ./backend/
COPY backend/public/ ./public/

RUN go mod download
RUN go build -o main ./backend/cmd/app

EXPOSE 8080
CMD ["./main"]