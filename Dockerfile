# Stage 1: Build the Go application
FROM golang:1.23.0-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application and output a binary named 'myShoppingApp'
RUN go build -o myShoppingApp

# Stage 2: Create a lightweight container to run the application
FROM alpine:latest

# Set the working directory in the new container
WORKDIR /root/

# Copy the 'myShoppingApp' binary from the 'build' stage into the new container
COPY --from=build /app/myShoppingApp .

# Copy the views and internal directories from the 'build' stage
COPY --from=build /app/views /root/views
COPY --from=build /app/internal /root/internal

# Expose port 8080 to the host
EXPOSE 8080

# Command to run when the container starts
CMD ["./myShoppingApp"]
