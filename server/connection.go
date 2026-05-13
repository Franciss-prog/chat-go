package main

// imports
import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Pool, error) {
	// validate the load of godotenv
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	// get the database url from env
	dbURL := os.Getenv("DATABASE_URL")

	// return the connection
	return pgxpool.New(context.Background(), dbURL)
}
