package messages

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
	// Connect to Redis without a password
	rdb = redis.NewClient(&redis.Options{
		Addr: "host.docker.internal:6379", // Update with your Redis server address
		DB:   0,                           // Default DB
	})

	// Check if the connection to Redis is successful
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
		panic(err)
	}
}

// Message represents a data structure for a message
type Message struct {
	Category       string `json:"category"`
	MessagePayload string `json:"messagePayload"`
}

// HandleMessage handles the "/handleMessage" endpoint
func HandleMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message

	// Parse JSON request body
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Check if message payload is not empty
	if msg.MessagePayload != "" {
		// Store message in Redis
		err := rdb.Set(context.Background(), msg.Category, msg.MessagePayload, 0).Err()
		if err != nil {
			http.Error(w, "Failed to store message in Redis", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Message handled successfully")
}

// GetAllMessages handles the "/getAllMessages" endpoint
func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	// Get all keys from Redis
	keys, err := rdb.Keys(context.Background(), "*").Result()
	if err != nil {
		http.Error(w, "Failed to retrieve messages from Redis", http.StatusInternalServerError)
		return
	}

	var messages []Message

	// Iterate over keys and retrieve messages
	for _, key := range keys {
		payload, err := rdb.Get(context.Background(), key).Result()
		if err != nil {
			log.Printf("Failed to retrieve message for key %s: %v", key, err)
			continue
		}

		// Create Message struct from retrieved data
		message := Message{
			Category:       key,
			MessagePayload: payload,
		}

		messages = append(messages, message)
	}

	// Respond with the list of messages as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

// ClearAllMessages handles the "/clearAllMessages" endpoint
func ClearAllMessages(w http.ResponseWriter, r *http.Request) {
	// Delete all keys from Redis
	err := rdb.FlushDB(context.Background()).Err()
	if err != nil {
		http.Error(w, "Failed to clear messages in Redis", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "All messages cleared successfully")
}
