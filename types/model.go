package types

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Item struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Category    uint   `json:"category"`
	Description string `json:"description"`
}

type Category struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
