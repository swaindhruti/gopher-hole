package repository

import (
	"blog-app/internal/models"
	"database/sql"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create inserts a new user into the database
func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (username, full_name, email, password_hash, role, bio, avatar_url, created_at, updated_at, is_active)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	now := time.Now()
	return r.db.QueryRow(
		query,
		user.Username,
		user.FullName,
		user.Email,
		user.PasswordHash,
		user.Role,
		user.Bio,
		user.AvatarURL,
		now,
		now,
		true,
	).Scan(&user.ID)
}

// GetByID retrieves a user by their ID
func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	query := `SELECT id, username, full_name, email, role, password_hash, bio, avatar_url, created_at, updated_at, is_active
			  FROM users WHERE id = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.FullName,
		&user.Email,
		&user.Role,
		&user.PasswordHash,
		&user.Bio,
		&user.AvatarURL,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.IsActive,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAll retrieves all users
func (r *UserRepository) GetAll() ([]*models.User, error) {
	query := `SELECT id, username, full_name, email, role, bio, avatar_url, created_at, updated_at, is_active
			  FROM users ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.FullName,
			&user.Email,
			&user.Role,
			&user.Bio,
			&user.AvatarURL,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.IsActive,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

// Update updates an existing user
func (r *UserRepository) Update(user *models.User) error {
	query := `UPDATE users SET full_name = $1, email = $2, role = $3, bio = $4, avatar_url = $5, updated_at = $6, is_active = $7
			  WHERE id = $8`

	_, err := r.db.Exec(
		query,
		user.FullName,
		user.Email,
		user.Role,
		user.Bio,
		user.AvatarURL,
		time.Now(),
		user.IsActive,
		user.ID,
	)
	return err
}

// Delete deletes a user by their ID
func (r *UserRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
