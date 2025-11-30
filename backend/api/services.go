package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jamalfox85/perfume-app/backend/data"
	"github.com/joho/godotenv" // Import the godotenv package
)

func NewApplication() *Application {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create DB Instance and table repositories
	db := newDBPool()

	profiles := data.NewProfileRepository(db)
	perfumes := data.NewPerfumeRepository(db)




	return &Application{
		DB:         db,
		Profiles:   profiles,
		Perfumes: perfumes,
	}
}

func newDBPool() *pgxpool.Pool {
	ctx := context.Background()
	dbURL := os.Getenv("DB_CONNECTION_STRING")

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Failed to create DB pool: %v", err)
	}

	// Test the connection
	var name string
	err = pool.QueryRow(ctx, "SELECT name FROM perfumes LIMIT 1").Scan(&name)
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}
	fmt.Println("First perfume name:", name)

	fmt.Println("Connected to the database successfully")
	return pool
}