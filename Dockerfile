FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd

FROM alpine:latest
WORKDIR /root/
# Копируем бинарник + добавляем зависимости
COPY --from=builder /app/main .
RUN apk add --no-cache libc6-compat  

EXPOSE 8000
CMD ["./main"]