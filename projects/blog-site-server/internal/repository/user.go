package repository

import (
	"blog-app/internal/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, email, password_hash, created_at, updated_at)
			  VALUES (?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, user.Username, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt)
	return err
}

// GetUserByID retrieves a user by their ID
func (r *UserRepository) GetUserByID(id int64) (*models.User, error) {
	query := `SELECT id, username, full_name, email, role, bio, avatar_url, created_at, updated_at, is_active
			  FROM users WHERE id = ?`
	row := r.DB.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Email, &user.Role, &user.Bio, &user.AvatarURL, &user.CreatedAt, &user.UpdatedAt, &user.IsActive)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates an existing user
func (r *UserRepository) UpdateUser(user *models.User) error {
	query := `UPDATE users SET full_name = ?, email = ?, role = ?, bio = ?, avatar_url = ?, updated_at = ?, is_active = ?
			  WHERE id = ?`
	_, err := r.DB.Exec(query, user.FullName, user.Email, user.Role, user.Bio, user.AvatarURL, user.UpdatedAt, user.IsActive, user.ID)
	return err
}

// DeleteUser deletes a user by their ID
func (r *UserRepository) DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}
