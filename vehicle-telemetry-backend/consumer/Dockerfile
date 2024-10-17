FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o consumer consumer.go
CMD ["./consumer"]