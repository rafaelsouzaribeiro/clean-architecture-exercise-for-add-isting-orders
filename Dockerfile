FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/ordersystem


FROM alpine:latest

RUN apk add --no-cache \
    ca-certificates \
    curl \
    make \
    bash \
    netcat-openbsd


RUN wget -O /tmp/migrate.tar.gz \
    https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz && \
    tar -xzf /tmp/migrate.tar.gz -C /tmp && \
    mv /tmp/migrate /usr/local/bin/migrate && \
    chmod +x /usr/local/bin/migrate

WORKDIR /app

COPY --from=builder /app/app .


COPY --from=builder /app/cmd/ordersystem/.env .env

COPY --from=builder /app/sql ./sql
COPY --from=builder /app/Makefile .
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

# Web server
EXPOSE 8000
# gRPC
EXPOSE 50051
# GraphQL
EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]