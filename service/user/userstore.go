package user

import (
	"database/sql"
	"fmt"
	"strings"

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
			return nil, fmt.Errorf("user not found")
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

func (s *Store) UpdateUser(user *types.User) (*types.User, error) {

	query := "UPDATE users SET "
	var args []any
	var sets []string

	if user.UserName != "" {
		sets = append(sets, "username = ?")
		args = append(args, user.UserName)
	}
	if user.FirstName != "" {
		sets = append(sets, "first_name = ?")
		args = append(args, user.FirstName)
	}
	if user.LastName != "" {
		sets = append(sets, "last_name = ?")
		args = append(args, user.LastName)
	}
	if user.Email != "" {
		sets = append(sets, "email = ?")
		args = append(args, user.Email)
	}
	if user.Password != "" {
		sets = append(sets, "password = ?")
		args = append(args, user.Password)
	}
	if user.Status != "" {
		sets = append(sets, "status = ?")
		args = append(args, user.Status)
	}

	if len(sets) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(sets, ", ") + " WHERE id = ?"
	args = append(args, user.ID)

	_, err := s.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return s.GetUserById(user.ID)
}

func (s *Store) ExcistsByUsernameAndEmail(username string, email string) (bool, error) {
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE username = ? OR email = ?)"
	var exist bool
	err := s.db.QueryRow(query, username, email).Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil

}
