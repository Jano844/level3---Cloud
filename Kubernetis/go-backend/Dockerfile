
# Stage 1: Build the Go binary
FROM golang:1.24 AS builder

WORKDIR /app
COPY main.go .
# ALl used Files
RUN go mod init myapp
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /root/
EXPOSE 8080
COPY --from=builder /app/main .

CMD ["./main"]

# # For creating image on mac
# docker buildx build --platform linux/amd64 -t go-backend .
