# Build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Run stage
FROM alpine:latest
# RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /app/main /app/app.env /
EXPOSE 8080
ENTRYPOINT ["/main"]
