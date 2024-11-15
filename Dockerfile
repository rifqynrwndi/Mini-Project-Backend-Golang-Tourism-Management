FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .env
EXPOSE 8000
CMD ["./main"]
