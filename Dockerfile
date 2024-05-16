# Use the official Golang image as the base image
FROM golang:1.20 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go binary
RUN go build -o main main.go

# Use an official Python runtime as the base image
FROM python:3.10-slim

# Set the working directory
WORKDIR /app

# Copy the requirements file
COPY requirements.txt .

# Install the Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the source code to the container
COPY . .

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8800

# Command to run the Gunicorn server
CMD ["gunicorn", "-w", "4", "-b", "0.0.0.0:8800", "umlgo:app"]
