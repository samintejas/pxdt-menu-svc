package store

import "projectx.io/drivethru/types"

type UserStore interface {
	GetUserByEmail(email string) (*types.User, error)
	CreateUser(user *User) error
	GetUserById(id uint) (*User, error)
}
