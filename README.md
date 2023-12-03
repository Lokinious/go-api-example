# go-api-example
practice go to improve on my own understanding

A simple Go application that interacts with Redis, designed for demonstration and learning purposes.

## Features

- Accepts messages with a category and payload.
- Stores non-empty messages in a Redis cache.
- Provides API routes to retrieve all messages and clear the cache.

## Prerequisites

- Go installed on your local machine.
- Docker installed on your local machine.
- Minikube installed for Kubernetes development.

## Getting Started

### Local Development

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/your-go-app.git
   cd your-go-app
   ```

2. Build and run the Go app locally:

   ```bash
   go run main.go
   ```

3. Access the app at [http://localhost:8080](http://localhost:8080).

### Docker Development

1. Build the Docker image:

   ```bash
   docker build -t yourusername/your-go-app:latest .
   ```

2. Run the Docker container:

   ```bash
   docker run -p 8080:8080 -e REDIS_ADDR=host.docker.internal:6379 yourusername/your-go-app:latest
   ```

3. Access the app at [http://localhost:8080](http://localhost:8080).

### Kubernetes Deployment (Minikube)

1. Start Minikube:

   ```bash
   minikube start
   ```

2. Deploy the Redis service:

   ```bash
   kubectl apply -f path/to/redis-deployment.yaml
   ```

3. Deploy the Go app:

   ```bash
   kubectl apply -f path/to/go-app-deployment.yaml
   ```

4. Access the Go app using the NodePort assigned:

   ```bash
   minikube service yourapp-service
   ```

## API Routes

- `POST /api/message`: Send a message to the app.
  ```json
  {
    "category": "your-category",
    "messagePayload": "your-payload"
  }
  ```

- `GET /api/messages`: Retrieve all messages stored in the cache.

- `DELETE /api/clear`: Clear all messages from the cache.

```

Replace placeholders like `yourusername`, `your-go-app`, and `your-category` with your actual details. This README provides a quick guide on local development, Docker usage, and Kubernetes deployment. Make sure to include details specific to your application and any additional instructions or information you want to convey.