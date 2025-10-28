package repository

import (
	"blog-app/internal/models"
	"database/sql"
	"time"
)

type BlogRepository struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

// Create inserts a new blog post into the database
func (r *BlogRepository) Create(blog *models.Blog) error {
	query := `INSERT INTO blogs (title, content, cover_image, author_id, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	now := time.Now()
	return r.db.QueryRow(
		query,
		blog.Title,
		blog.Content,
		blog.CoverImage,
		blog.AuthorID,
		now,
		now,
	).Scan(&blog.ID)
}

// GetByID retrieves a blog post by its ID
func (r *BlogRepository) GetByID(id int64) (*models.Blog, error) {
	query := `SELECT id, title, content, cover_image, author_id, created_at, updated_at
			  FROM blogs WHERE id = $1`

	blog := &models.Blog{}
	err := r.db.QueryRow(query, id).Scan(
		&blog.ID,
		&blog.Title,
		&blog.Content,
		&blog.CoverImage,
		&blog.AuthorID,
		&blog.CreatedAt,
		&blog.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

// GetAll retrieves all blog posts
func (r *BlogRepository) GetAll() ([]*models.Blog, error) {
	query := `SELECT id, title, content, cover_image, author_id, created_at, updated_at
			  FROM blogs ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []*models.Blog
	for rows.Next() {
		blog := &models.Blog{}
		err := rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Content,
			&blog.CoverImage,
			&blog.AuthorID,
			&blog.CreatedAt,
			&blog.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, rows.Err()
}

// Update updates an existing blog post
func (r *BlogRepository) Update(blog *models.Blog) error {
	query := `UPDATE blogs SET title = $1, content = $2, cover_image = $3, updated_at = $4
			  WHERE id = $5`

	_, err := r.db.Exec(
		query,
		blog.Title,
		blog.Content,
		blog.CoverImage,
		time.Now(),
		blog.ID,
	)
	return err
}

// Delete deletes a blog post by its ID
func (r *BlogRepository) Delete(id int64) error {
	query := `DELETE FROM blogs WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
