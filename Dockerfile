# Use an official Golang image as the base image
FROM golang:1.22-alpine3.20

# Set the environment variable for the timezone
ENV TZ="Asia/Tashkent"

# Create and set the working directory
RUN mkdir /app
WORKDIR /app

# Copy go.mod and go.sum files separately to leverage Docker caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy && go mod vendor

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o app ./cmd/app/main.go

# Command to run the executable
CMD ["./app"]
