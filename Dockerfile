# Use a base image with Go pre-installed
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the application, assuming your main.go or its package is in the './cmd' directory
RUN go build -o /app/cmd/myapp ./cmd

# Command to run the executable, adjusted to point to the correct location
CMD ["/app/cmd/myapp"]
