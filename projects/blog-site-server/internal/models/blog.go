package models

import (
	"time"
)

// Blog model
type Blog struct {
	ID         int64      `json:"id"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	CoverImage string     `json:"cover_image"`
	AuthorID   int64      `json:"author_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
