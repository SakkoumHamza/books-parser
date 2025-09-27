FROM golang:1.25.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Set environment variables
ENV MONGO_URI=mongodb://localhost:27017
ENV MONGO_DATABASE=local
ENV RABBITMQ_URI=amqp://guest:guest@localhost:5672/

# Build the binary
RUN go build -o app main.go

# Start the app when container runs
CMD ["./app"]
