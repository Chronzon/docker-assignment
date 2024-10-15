# Step 1: Use an official Golang runtime as a parent image
FROM golang:1.20

# Step 2: Set the current working directory inside the container
WORKDIR /app

# Step 3: Copy the Go modules file (if it exists) and go source code
COPY main.go .

# Step 4: Build the Go application
RUN go mod init example.com/app || true
RUN go build -o main .

# Step 5: Expose port 8080
# Port yang di expose ini nanti dimasukin ke containerPort di pod-assignment.yaml
EXPOSE 8080

# Step 6: Command to run the executable
CMD ["/app/main"]
