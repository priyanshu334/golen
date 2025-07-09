# Start from the official Go image
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum if you have them (optional for dependency resolution)
COPY go.mod  ./
RUN go mod download

# Copy the rest of your source code
COPY . .

# Build the Go binary
RUN go build -o main .

# Run the binary
CMD ["./main"]
