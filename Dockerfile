FROM golang:1.25-alpine AS go-builder
WORKDIR /app
RUN apk add --no-cache gcc musl-dev
COPY web/ ./
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=go-builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main", "-prod"]
