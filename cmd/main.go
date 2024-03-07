package main

import (
	"GopherGate/pkg/handler"
	"GopherGate/pkg/router"
)

func main() {
	// Start the server
	startServer()
}

// startServer initializes the router and starts the server on port 1323.
func startServer() {
	// Create a new handler for users
	usersHandler := handler.NewUsersHandler()

	// Initialize the router
	r := router.NewRouter(*usersHandler).InitRouter()

	// Start the server on port 1323
	err := r.WebApiFramework.Listen(":1323")
	if err != nil {
		panic(err)
	}
}
