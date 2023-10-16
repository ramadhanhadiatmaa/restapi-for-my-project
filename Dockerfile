# Stage 1: Build the application
FROM golang:1.21.1 AS build

WORKDIR /app

# Copy Go module and source code
COPY go.mod go.sum ./
COPY . .

# Download module dependencies
RUN go mod download

# Build the Go application
RUN CGO_ENABLED=0 go build -o apicsmfib

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the previous stage
COPY --from=build /app/apicsmfib .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./apicsmfib"]

