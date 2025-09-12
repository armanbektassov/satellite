# Stage 1: Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Сначала только go.mod + go.sum (чтобы кешировать зависимости)
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарь из cmd/main.go
RUN go build -o app ./cmd/main.go

# Stage 2: Run (минимальный образ)
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
