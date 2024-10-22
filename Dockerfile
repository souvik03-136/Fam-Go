# Use Golang base image
FROM golang:1.19-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy Go modules manifest and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code to the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/api

# Expose the port that the app will run on
EXPOSE 8080

# Command to run the app
CMD ["./main"]
