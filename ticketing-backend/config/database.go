package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

// ConnectToDB establishes and returns a connection to the database
func ConnectToDB() *gorm.DB {
	log.Println("Establishing database connection...")
	dsn := os.Getenv("DB_DSN")

	// Check if the DSN environment variable is set
	if dsn == "" {
		log.Println("Error: Environment variable DB_DSN is not set")
		log.Println("Please set the DB_DSN environment variable in your .env file and try again.")
		os.Exit(1)
	}

	var db *gorm.DB
	var err error

	// Attempt to connect to the database with retries
	for attempts := 1; attempts <= 3; attempts++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Database connection established successfully.")
			return db
		}

		log.Printf("Attempt %d: Failed to connect to the database: %v\n", attempts, err)

		// If it's not the last attempt, wait before retrying
		if attempts < 3 {
			log.Println("Retrying in 2 seconds...")
			time.Sleep(2 * time.Second)
		}
	}

	// If all attempts fail, log the final error and exit
	log.Fatalf("Final attempt failed: Unable to establish database connection after multiple retries: %v", err)
	return nil
}
