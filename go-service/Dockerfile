FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go
EXPOSE 8080
CMD ["./server"]
