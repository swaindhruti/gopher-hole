package repository

import (
	"blog-app/internal/models"
	"database/sql"
)

type CommentRepository struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

// CreateComment inserts a new comment into the database
func (r *CommentRepository) CreateComment(comment *models.Comment) error {
	query := `INSERT INTO comments (blog_id, user_id, content, created_at, updated_at)
			  VALUES (?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, comment.PostID, comment.UserID, comment.Content, comment.CreatedAt, comment.UpdatedAt)
	return err
}

// GetCommentsByBlogID retrieves all comments for a specific blog post
func (r *CommentRepository) GetCommentsByBlogID(blogID int64) ([]*models.Comment, error) {
	query := `SELECT id, blog_id, user_id, content, created_at, updated_at
			  FROM comments WHERE blog_id = ?`
	rows, err := r.DB.Query(query, blogID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		comment := &models.Comment{}
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

// DeleteComment deletes a comment by its ID
func (r *CommentRepository) DeleteComment(id int64) error {
	query := `DELETE FROM comments WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

// UpdateComment updates an existing comment
func (r *CommentRepository) UpdateComment(comment *models.Comment) error {
	query := `UPDATE comments SET content = ?, updated_at = ?
			  WHERE id = ?`
	_, err := r.DB.Exec(query, comment.Content, comment.UpdatedAt, comment.ID)
	return err
}
