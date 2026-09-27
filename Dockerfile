# Use the newer 'bookworm' (Debian 12) environment
FROM golang:1.27-bookworm

# Install Python and pip
RUN apt-get update && apt-get install -y python3 python3-pip

# Set up the working directory
WORKDIR /app
COPY . .

# Build the Go web server
RUN go build -o server .

# Expose the port and run the server
EXPOSE 3000
CMD ["./server"]