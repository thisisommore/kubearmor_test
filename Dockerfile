# Use the official Golang base image
FROM golang:1.19-alpine as builder

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files to cache the dependencies
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN GOOS=linux go build -o kubearmor_test .

# Use the scratch image for the final output
FROM busybox

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/kubearmor_test .

# Run the binary
ENTRYPOINT ["./kubearmor_test"]
