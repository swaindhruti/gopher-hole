package repository

import (
	"blog-app/internal/models"
	"database/sql"
)

type BlogRepository struct {
	DB *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{DB: db}
}

// CreateBlog inserts a new blog post into the database
func (r *BlogRepository) CreateBlog(blog *models.Blog) error {
	query := `INSERT INTO blogs (title, content, cover_image, author_id, created_at, updated_at)
			  VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, blog.Title, blog.Content, blog.CoverImage, blog.AuthorID, blog.CreatedAt, blog.UpdatedAt)
	return err
}

// GetBlogByID retrieves a blog post by its ID
func (r *BlogRepository) GetBlogByID(id int64) (*models.Blog, error) {
	query := `SELECT id, title, content, cover_image, author_id, created_at, updated_at
			  FROM blogs WHERE id = ?`
	row := r.DB.QueryRow(query, id)

	blog := &models.Blog{}
	err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.CoverImage, &blog.AuthorID, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

// UpdateBlog updates an existing blog post
func (r *BlogRepository) UpdateBlog(blog *models.Blog) error {
	query := `UPDATE blogs SET title = ?, content = ?, cover_image = ?, updated_at = ?
			  WHERE id = ?`
	_, err := r.DB.Exec(query, blog.Title, blog.Content, blog.CoverImage, blog.UpdatedAt, blog.ID)
	return err
}

// DeleteBlog deletes a blog post by its ID
func (r *BlogRepository) DeleteBlog(id int64) error {
	query := `DELETE FROM blogs WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}
