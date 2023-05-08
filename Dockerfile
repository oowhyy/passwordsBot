FROM golang:1.20-alpine as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /main cmd/main/main.go
# Финальный этап, копируем собранное приложение
FROM alpine:latest
COPY --from=builder main /bin/main
COPY .env .
COPY config.yml .
ENTRYPOINT ["/bin/main"]