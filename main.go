package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/lokinious/go-api-example/routes"
)

var rdb *redis.Client

func init() {
	redisAddr := getEnv("REDIS_ADDR", "host.docker.internal:6379")

	// Connect to Redis without a password
	rdb = redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   0, // Default DB
	})

	// Check if the connection to Redis is successful
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
		os.Exit(1)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	// Check if Redis connection is successful before starting the server
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis connection is not established: %v. Exiting.", err)
		os.Exit(1)
	}

	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define routes
	routes.DefineAllRoutes(router)

	// Start the HTTP server
	http.Handle("/", router)
	fmt.Println("Server is running on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
