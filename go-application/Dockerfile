# Use the official Golang image as the base image
FROM golang:1.22.2-alpine3.19 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -C ./cmd/app

# Set the entry point for the container
ENTRYPOINT ["./cmd/app/app"]