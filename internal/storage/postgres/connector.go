package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func NewDb() (*sql.DB, error) {
	err := godotenv.Load("../../.env") // loads .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("✅ Connected to DB")
	return db, nil
}
