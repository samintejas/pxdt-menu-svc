package store

import "projectx.io/drivethru/types"

type UserStore interface {
	GetUserByEmail(email string) (*types.User, error)
	CreateUser(user *types.User) (uint, error)
	GetUserById(id uint) (*types.User, error)
}

type ItemStore interface {
	GetAllItems() (*types.Item, error)
	CreateItem(item *types.Item) (uint, error)
	GetItemById(id uint) (*types.Item, error)
}

type CategoryStore interface {
	GetCategory() (*types.Category, error)
	CreateCategory(category *types.Category) (uint, error)
	GetCategoryById(id uint) (*types.Category, error)
}
