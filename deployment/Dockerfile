FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o game ./app

# Setup the runtime container
FROM alpine:latest  
# Add ca-certificates in case you need HTTPS
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/game .
COPY --from=builder /app/config ./config
CMD ["./game"]

