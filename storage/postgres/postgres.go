package postgres

import (
	"database/sql"
	"fmt"
	"github.com/AlexMykhailov1/ImageAPI/config"
	_ "github.com/lib/pq"
)

// NewPostgres creates connection to new postgres db using config file.
// Connection is not being closed
func NewPostgres(cfg *config.Config) (*sql.DB, error) {
	// Set data source name from .env file
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSL)

	// Open connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Ping the database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Postgres connect successful!")

	return db, nil
}
