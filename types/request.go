package types

type RegisterUser struct {
	UserName  string `json:"userName" validate:"required,min=2,max=32,alphanum"`
	FirstName string `json:"firstName" validate:"required,min=1,max=32,alpha"`
	LastName  string `json:"lastName" validate:"required,min=1,max=32,alpha"`
	Email     string `json:"email" validate:"required,email,max=255"`
	Password  string `json:"password" validate:"required,min=8,max=64,containsany=!@#$%^&*"`
	Status    string `json:"status" validate:"required,oneof=active inactive"`
}

type UpdateUser struct {
	UserName  string `json:"userName,omitempty" validate:"omitempty,min=2,max=32,alphanum"`
	FirstName string `json:"firstName,omitempty" validate:"omitempty,min=1,max=32,alpha"`
	LastName  string `json:"lastName,omitempty" validate:"omitempty,min=1,max=32,alpha"`
	Email     string `json:"email,omitempty" validate:"omitempty,email,max=255"`
	Password  string `json:"password,omitempty" validate:"omitempty,min=8,max=64,containsany=!@#$%^&*"`
	Status    string `json:"status,omitempty" validate:"omitempty,oneof=active inactive"`
}

type CreateItem struct {
	Name        string `json:"name"`
	Category    uint   `json:"category"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateItem struct {
	Name        string `json:"name,omitempty"`
	Category    uint   `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}

type UpdateCategory struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}
