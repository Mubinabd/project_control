FROM golang:1.22.5 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o project ./cmd
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/project .
COPY .env .env 
RUN chmod +x project
EXPOSE 8080
CMD ["./project"]