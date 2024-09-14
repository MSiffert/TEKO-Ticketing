package main

import (
	"example-rest-api/app/router"
	"example-rest-api/config"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	log.Println("Loading environment variables...")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("Environment variables loaded successfully.")

	log.Println("Initializing logging configuration...")
	config.InitLog()
	log.Println("Logging configuration initialized.")
}

func main() {
	log.Println("Starting the application...")

	// Retrieve the port from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("Environment variable PORT is not set")
	}
	log.Printf("Port set to %s", port)

	// Initialize the application dependencies
	log.Println("Initializing application dependencies...")
	init := config.Init()
	log.Println("Application dependencies initialized successfully.")

	// Set up the router with initialized dependencies
	log.Println("Setting up the router...")
	app := router.Init(init)
	log.Println("Router setup completed.")

	// Run the application on the specified port
	log.Printf("Running the application on port %s...", port)
	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Error running the application: %v", err)
	}
	log.Println("Application is running.")
}
