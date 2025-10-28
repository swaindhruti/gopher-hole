package repository

import (
	"blog-app/internal/models"
	"database/sql"
	"time"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

// Create inserts a new comment into the database
func (r *CommentRepository) Create(comment *models.Comment) error {
	query := `INSERT INTO comments (post_id, user_id, content, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`

	now := time.Now()
	return r.db.QueryRow(
		query,
		comment.PostID,
		comment.UserID,
		comment.Content,
		now,
		now,
	).Scan(&comment.ID)
}

// GetByID retrieves a comment by its ID
func (r *CommentRepository) GetByID(id int64) (*models.Comment, error) {
	query := `SELECT id, post_id, user_id, content, created_at, updated_at
			  FROM comments WHERE id = $1`

	comment := &models.Comment{}
	err := r.db.QueryRow(query, id).Scan(
		&comment.ID,
		&comment.PostID,
		&comment.UserID,
		&comment.Content,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// GetByBlogID retrieves all comments for a specific blog post
func (r *CommentRepository) GetByBlogID(blogID int64) ([]*models.Comment, error) {
	query := `SELECT id, post_id, user_id, content, created_at, updated_at
			  FROM comments WHERE post_id = $1 ORDER BY created_at DESC`

	rows, err := r.db.Query(query, blogID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		comment := &models.Comment{}
		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, rows.Err()
}

// Update updates an existing comment
func (r *CommentRepository) Update(comment *models.Comment) error {
	query := `UPDATE comments SET content = $1, updated_at = $2 WHERE id = $3`

	_, err := r.db.Exec(query, comment.Content, time.Now(), comment.ID)
	return err
}

// Delete deletes a comment by its ID
func (r *CommentRepository) Delete(id int64) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
