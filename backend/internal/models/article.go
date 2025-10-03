package models

import "time"

type Article struct {
	ID          int64     `json:"id" db:"id"`
	FeedID      int64     `json:"feed_id" db:"feed_id"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	URL         string    `json:"url" db:"url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
}
