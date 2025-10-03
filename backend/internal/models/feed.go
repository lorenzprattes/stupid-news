package models

import (
	"time"
)

type Feed struct {
	ID           int64     `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	URL          string    `json:"url" db:"url"`
	SiteURL      string    `json:"site_url" db:"site_url"`
	Description  string    `json:"description" db:"description"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	ErrorCount   int       `json:"error_count" db:"error_count"`
	FetchCount   int       `json:"fetch_count" db:"fetch_count"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	LastFetched  time.Time `json:"last_fetched" db:"last_fetched"`
}
