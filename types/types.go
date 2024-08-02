package types

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
}

func GetUserByEmail(email string) (*User, error) {
	return nil, nil
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
