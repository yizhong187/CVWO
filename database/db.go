package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// initDB initializes the database connection
func InitDB() {

	// Load environment variblaes from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve database connection parameters from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbSslmode := os.Getenv("DB_SSLMODE")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Open new database connection and assign it to the global 'db' variable
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		dbHost, dbPort, dbUser, dbName, dbSslmode, dbPassword)
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
}

// GetDB returns the global database connection.
// This can be used to perform database operations using the established connection.
func GetDB() *sql.DB {
	return db
}
