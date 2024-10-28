# Build stage
FROM golang:1.23.2-alpine3.20 AS builder

# Set working directory
WORKDIR /app

# Copy source code
COPY . . 

# Install dependencies
RUN go mod download

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/samsamoohooh ./cmd/app/app.go

# Final stage
FROM alpine:3.20 As Final
LABEL maintainer="fullgukbap"

# Set working directory
WORKDIR /app

# Copy binary
COPY --from=builder /app/bin/samsamoohooh ./

EXPOSE 8080

ENTRYPOINT [ "./samsamoohooh" ]