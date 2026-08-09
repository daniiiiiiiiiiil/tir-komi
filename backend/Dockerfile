FROM golang:1.26.4-alpine AS builder

WORKDIR /app

COPY backend/go.mod ./
RUN go mod download

COPY backend/ .

RUN go build -o main ./cmd/app

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/main .
COPY --from=builder /app/public ./public

EXPOSE 8080

CMD ["./main"]