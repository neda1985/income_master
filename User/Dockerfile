# Use the official Golang image as the base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the application source code to the container
COPY . .

# Build the application binary inside the container
RUN go build -o myapp

# Expose the port that your application listens on
EXPOSE 8080

# Start the application when the container starts
CMD ["./myapp"]
