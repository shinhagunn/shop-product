FROM golang:1.16-alpine3.15 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build main.go

FROM alpine:3.15 AS production

WORKDIR /app

COPY --from=builder /app/main /app/main

CMD ["./main"]

