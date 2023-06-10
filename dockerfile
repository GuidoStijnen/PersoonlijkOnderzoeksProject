# Use the official Golang Docker image as the base
FROM golang:1.20 AS builder

# Set the working directory to the location of the Go code
WORKDIR /app

# Copy the Go module files and the Go code to the working directory
COPY go.mod go.sum ./
COPY Main.go ./

# Download and install the Go modules
RUN go mod download

# Build the Go applicatie
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a minimal base image to reduce the size of the final Docker image
FROM alpine:latest

# Set the working directory 
WORKDIR /app/mainv2

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main /app/mainv2/main

# Copy the HTML templates to the working directory inside the container
COPY templates/*.html ./templates/

# Copy the config.yaml file to the working directory inside the container
COPY config.yaml ./config.yaml

# Expose port 8080 for accessing the application
EXPOSE 8080

# Run the Go binary
CMD ["./main"]
