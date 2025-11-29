# Build stage for Tailwind CSS
FROM node:20-alpine AS node-builder

WORKDIR /app

# Copy package files and install dependencies
COPY web/package*.json ./
RUN npm ci

# Copy static assets and build CSS
COPY web/ui ./ui
COPY web/tailwind.config.js ./
RUN npm run build:css

# Go build stage
FROM golang:1.25-alpine AS go-builder
WORKDIR /app
RUN apk add --no-cache gcc musl-dev
COPY web/ ./

# Copy built CSS from node-builder stage
COPY --from=node-builder /app/ui ./ui

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main .

# Final stage
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=go-builder /app/main /app/main
COPY --from=go-builder /app/ui ./ui

EXPOSE 8080

CMD ["/app/main", "-prod"]
