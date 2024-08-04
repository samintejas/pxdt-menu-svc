package user

import (
	"database/sql"
	"fmt"

	"projectx.io/drivethru/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	query := `
        SELECT id, username, email, password, first_name, last_name, status, created_at, updated_at
        FROM users
        WHERE email = ?
    `
	return s.getUserByQuery(query, email)
}

func (s *Store) GetUserByEmailAndStatus(email string, status string) (*types.User, error) {
	query := `
        SELECT id, username, email, password, first_name, last_name, status, created_at, updated_at
        FROM users
        WHERE email = ?
    `
	return s.getUserByQuery(query, email)
}

func (s *Store) GetUserById(id uint) (*types.User, error) {
	query := `
        SELECT id, username, email, password, first_name, last_name, status, created_at, updated_at
        FROM users
        WHERE id = ?
    `
	return s.getUserByQuery(query, id)
}

func (s *Store) GetUserByIdAndStatus(email string, status string) (*types.User, error) {
	query := `
        SELECT id, username, email, password, first_name, last_name, status, created_at, updated_at
        FROM users
        WHERE id = ? AND status = ?
    `
	return s.getUserByQuery(query, email, status)
}

func (s *Store) getUserByQuery(query string, args ...any) (*types.User, error) {

	var user types.User
	err := s.db.QueryRow(query, args...).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return &user, nil
}

func (s *Store) CreateUser(user *types.User) (uint, error) {

	query := "INSERT INTO users (username,first_name,last_name,email,password,status) values (?,?,?,?,?,?)"
	result, err := s.db.Exec(query, user.UserName, user.FirstName, user.LastName, user.Email, user.Password, user.Status)

	if err != nil {
		return 0, err
	}

	lastInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint(lastInserted), nil

}
