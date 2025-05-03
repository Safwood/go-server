FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd  # Путь к main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# Порт вашего сервиса
EXPOSE 8000
CMD ["./main"]