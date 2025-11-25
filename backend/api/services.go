package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv" // Import the godotenv package
)

func NewApplication() *Application {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create DB Instance and table repositories
	db := newDB()
	defer db.Close(context.Background())

	// messages := data.NewMessageRepository(db)
	// customers := data.NewCustomerRepository(db)


	return &Application{
		// Messages: messages,
		// Customers: customers,
	}
}

func newDB() *pgx.Conn {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close(ctx)

	// Example query to test connection
	var name string
    err = db.QueryRow(ctx, "SELECT name FROM perfumes LIMIT 1").Scan(&name)
    if err != nil {
        log.Fatalf("Query failed: %v\n", err)
    }
	fmt.Println("First perfume name:", name)

	fmt.Println("Connected to the database successfully")
	return db;
}