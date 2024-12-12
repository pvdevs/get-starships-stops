FROM golang:1.23-alpine as base
WORKDIR /app

# Development stage
FROM base as development
# Install air for hot reloading in development
RUN go install github.com/air-verse/air@latest
COPY go.* ./
RUN go mod download
COPY . .
EXPOSE 8080
CMD ["air", "-c", ".air.toml"]

# Production stage
FROM base as production
COPY go.* ./
RUN go mod download
COPY . .
# Build optimized binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o app ./cmd/app
CMD ["./app"]