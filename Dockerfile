FROM node:alpine AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build

FROM golang:1.23-alpine AS go-builder
WORKDIR /go/src/american-empire
RUN apk add --no-cache gcc musl-dev
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
COPY backend/ ./backend/
COPY cmd/ ./cmd/
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=go-builder /go/src/american-empire/main /app/main
COPY --from=frontend-builder /app/frontend/build /app/frontend/build

EXPOSE 3000

CMD ["/app/main"]
