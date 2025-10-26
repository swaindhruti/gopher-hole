package models

import (
	"database/sql"
	"time"
)

// Blog model
type Blog struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CoverImage string    `json:"cover_image"`
	AuthorID   int64     `json:"author_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Create a new blog post
func CreateBlog(db *sql.DB, blog *Blog) error {
	query := `INSERT INTO blogs (title, content, cover_image, author_id) 
	          VALUES ($1, $2, $3, $4) 
	          RETURNING id, title, created_at, updated_at`

	err := db.QueryRow(
		query, blog.Title, blog.Content, blog.CoverImage, blog.AuthorID,
	).Scan(&blog.ID, &blog.CreatedAt, &blog.UpdatedAt)

	return err
}

// Get blog post by ID
func GetBlogByID(db *sql.DB, id int64) (*Blog, error) {
	blog := &Blog{}
	query := `SELECT id, title, content, cover_image, author_id, created_at, updated_at 
	          FROM blogs WHERE id = $1`

	err := db.QueryRow(query, id).Scan(
		&blog.ID, &blog.Title, &blog.Content, &blog.CoverImage,
		&blog.AuthorID, &blog.CreatedAt, &blog.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return blog, nil
}

// Update blog post details
func UpdateBlog(db *sql.DB, blog *Blog) error {
	query := `UPDATE blogs SET title=$1, content=$2, cover_image=$3, updated_at=NOW() 
	          WHERE id=$4`

	_, err := db.Exec(
		query, blog.Title, blog.Content, blog.CoverImage, blog.ID,
	)

	return err
}

// Delete blog post
func DeleteBlog(db *sql.DB, id int64) error {
	query := `DELETE FROM blogs WHERE id=$1`
	_, err := db.Exec(query, id)
	return err
}
