package models

import (
	"database/sql"
	"time"
)

// User model
type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	FullName     string    `json:"name"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	PasswordHash string    `json:"-"`
	Bio          string    `json:"bio"`
	AvatarURL    string    `json:"avatar_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsActive     bool      `json:"is_active"`
}

// Create a new user
func CreateUser(db *sql.DB, user *User) error {
	query := `INSERT INTO users (username, email, password_hash, full_name, bio, avatar_url, role) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7) 
	          RETURNING id, created_at, updated_at, is_active`

	err := db.QueryRow(
		query, user.Username, user.Email, user.PasswordHash,
		user.FullName, user.Bio, user.AvatarURL, user.Role,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.IsActive)

	return err
}

// Get user by ID
func GetUserByID(db *sql.DB, id int64) (*User, error) {
	user := &User{}
	query := `SELECT id, username, email, full_name, bio, avatar_url, role, created_at, updated_at, is_active 
	          FROM users WHERE id = $1`

	err := db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.FullName,
		&user.Bio, &user.AvatarURL, &user.Role,
		&user.CreatedAt, &user.UpdatedAt, &user.IsActive,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update user details
func UpdateUser(db *sql.DB, user *User) error {

	query := `UPDATE users SET username=$1, email=$2, full_name=$3, bio=$4, avatar_url=$5, role=$6, updated_at=NOW() 
	          WHERE id=$7`

	_, err := db.Exec(
		query, user.Username, user.Email, user.FullName,
		user.Bio, user.AvatarURL, user.Role, user.ID,
	)

	return err
}

// State user as active or inactive
func SetUserActiveStatus(db *sql.DB, id int64, isActive bool) error {
	query := `UPDATE users SET is_active=$1, updated_at=NOW() WHERE id=$2`
	_, err := db.Exec(query, isActive, id)
	return err
}
