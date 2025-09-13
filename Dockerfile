FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/api

FROM alpine:latest
RUN apk add --no-cache ca-certificates postgresql-client
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 9000
CMD sh -c "until pg_isready -h $DB_HOST -p $DB_PORT -U $DB_USER; do echo waiting for database; sleep 2; done; ./main"