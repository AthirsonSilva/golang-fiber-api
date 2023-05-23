# Start with a base Golang image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port on which the API will run
EXPOSE 8000

# Set the command to run the executable when the container starts
CMD ["./main"]
