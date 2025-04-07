package users

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *User) error {
	query := `INSERT INTO users (id, username, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, user.ID, user.Username, user.Email, user.PasswordHash, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindAll() ([]User, error) {
	query := `SELECT id, username, email, password_hash, created_at FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) FindByID(id uuid.UUID) (*User, error) {
	query := `SELECT id, username, email, password_hash, created_at FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var user User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*User, error) {
	query := `SELECT id, username, email, password_hash, created_at FROM users WHERE username = $1`
	row := r.db.QueryRow(query, username)

	var user User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	query := `SELECT id, username, email, password_hash, created_at FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	var user User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *User) error {
	query := `UPDATE users SET username = $1, email = $2, password_hash = $3 WHERE id = $4`
	_, err := r.db.Exec(query, user.Username, user.Email, user.PasswordHash, user.ID)
	if err != nil {
		return err
	}
	return nil
}
