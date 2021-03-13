FROM golang:1.15 AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY ./cmd ./cmd
COPY ./internal ./internal

# RSA
COPY ./access-private.pem ./access-private.pem
COPY ./access-public.pem ./access-public.pem
COPY ./refresh-private.pem ./refresh-private.pem
COPY ./refresh-public.pem ./refresh-public.pem

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

FROM alpine:latest
COPY --from=builder /app .
CMD ["./main"]