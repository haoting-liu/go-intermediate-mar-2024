FROM golang:1.20-alpine3.19 AS builder

WORKDIR /go/app
COPY ./ .

RUN go mod download
RUN go build -o server

# app package in minimal OS
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /go/app/server .

CMD ["/app/server"]