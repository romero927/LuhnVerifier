# Start from the official golang image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod, go.sum (if it exists), and main.go
COPY go.mod go.sum* main.go ./

# Download the dependencies
RUN go mod download

# Build the application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 4000

# Run the binary
CMD ["./main"]
