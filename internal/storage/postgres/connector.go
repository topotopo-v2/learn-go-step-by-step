package postgres

import (
	"database/sql"
	"log/slog"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func NewDb(log *slog.Logger) (*sql.DB, error) {
	err := godotenv.Load(".env") // loads .env
	if err != nil {
		log.Warn("no .env file found, relying on environment variables")
	}

	dsn := os.Getenv("DB_URL")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Error("failed to open database connection", "error", err)
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Error("failed to ping database", "error", err)
		return nil, err
	}

	log.Info("Connected to DB")
	return db, nil
}
