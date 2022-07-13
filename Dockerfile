# Start from the lastest version of golang image
FROM golang:1.18-alpine

# Add maintainer information
LABEL maintainer="Quique <thangvv@ftech.ai>"

# Set current directory as working directory
WORKDIR /app

# Copy go modules dependencies required file
COPY go.mod .

# Copy go modules expected hash file
COPY go.sum .

# Download go modules dependencies
RUN go mod download

# Copy all the app soures 
COPY . .

# Set http port
ENV PORT 8080

# Build the app
RUN go build 

# Remove soure files
RUN find . -name "*.go" -type f -delete

# Make port 8080 available to the world
EXPOSE 8080

# Run app
CMD ["./PR_gin_g"]

