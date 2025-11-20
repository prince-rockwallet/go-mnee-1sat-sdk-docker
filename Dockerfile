FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -o mnee-server cmd/server/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/mnee-server .
COPY .env . 

EXPOSE 8080
CMD ["./mnee-server"]