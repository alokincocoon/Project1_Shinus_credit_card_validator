# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod .
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o creditcardvalidator

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./creditcardvalidator"]
