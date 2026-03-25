FROM golang:alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
# Копируем папку vendor (теперь зависимости локальные!)
COPY vendor ./vendor

# ВАЖНО: Вместо go mod download, мы сразу копируем исходники
COPY . .

# Собираем с флагом -mod=vendor
RUN go build -mod=vendor -o main ./cmd/api

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]