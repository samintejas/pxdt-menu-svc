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

	rows, err := s.db.Query("SELECT * FROM users WHERE email= ?", email)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {

		u := new(types.User)
		err = scanRowIntoUser(rows, u)
		if err != nil {
			return nil, err
		}

		if u.ID == 0 {
			return nil, fmt.Errorf("user not found")
		}
		return u, nil
	}
	return nil, nil
}

func scanRowIntoUser(rows *sql.Rows, user *types.User) error {

	return rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
	)

}

func (s *Store) GetUserById(id uint) (*types.User, error) {

	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {

		u := new(types.User)
		err = scanRowIntoUser(rows, u)
		if err != nil {
			return nil, err
		}

		if u.ID == 0 {
			return nil, fmt.Errorf("user not found")
		}
		return u, nil
	}
	return nil, nil
}

func (s *Store) CreateUser(user *types.User) (uint, error) {

	query := "INSERT INTO users (first_name,last_name,email,password) values (?,?,?,?)"
	result, err := s.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint(lastInserted), nil

}
