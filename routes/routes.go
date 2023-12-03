package routes

import (
	"github.com/gorilla/mux"
	"github.com/lokinious/go-api-example/messages"
)

// DefineAllRoutes sets up all routes for the application
func DefineAllRoutes(router *mux.Router) {
	// Define the route for handling messages
	router.HandleFunc("/handleMessage", messages.HandleMessage).Methods("POST")

	// Define the route for getting all messages
	router.HandleFunc("/getAllMessages", messages.GetAllMessages).Methods("GET")

	// Define the route for clearing all messages
	router.HandleFunc("/clearAllMessages", messages.ClearAllMessages).Methods("DELETE")
}
