package item

import (
	"database/sql"
	"time"

	"projectx.io/drivethru/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllItems() (*types.Item, error) {
	return &types.Item{
		ID:          0,
		Name:        "test",
		Category:    0,
		Description: "test desc",
		Status:      "Deactive",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (s *Store) CreateItem(item *types.Item) (uint, error) {

	query := "INSERT INTO items (name,category,description,status) values (?,?,?,?,?,?)"

	result, err := s.db.Exec(query, item.Name, item.Category, item.Description, item.Status)
	if err != nil {
		return 0, err
	}
	lastInserted, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint(lastInserted), nil
}

func (s *Store) GetItemById(id uint) (*types.Item, error) {
	return &types.Item{
		ID:          0,
		Name:        "test",
		Category:    0,
		Description: "test desc",
		Status:      "Deactive",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
