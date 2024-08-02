package user

import (
	"database/sql"
	"fmt"
	"log"

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

	u := new(types.User)

	for rows.Next() {
		err = scanRowIntoUser(rows, u)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
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

func (s *Store) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(user *types.User) error {

	query := "INSERT INTO users (first_name,last_name,email,password) values (?,?,?,?)"
	log.Println(user)
	result, err := s.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return err
	}

	lastInserted, err := result.LastInsertId()
	if err != nil {
		return err
	}

	log.Printf("Records inserted %d \n", lastInserted)

	return nil

}
