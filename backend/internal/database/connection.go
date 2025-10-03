package database

import (
	"database/sql"
	"fmt"
	"stupid-news/backend/internal/config"
	"time"

	_ "github.com/lib/pq"
)

func ConnectionString(c *config.DatabaseConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode,
	)
}

func NewConnection(dbConfig *config.DatabaseConfig) (*sql.DB, error) {
	connectionString := ConnectionString(dbConfig)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)
	return db, nil
}
