package item

type MenuItem struct {
	ID           uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string  `gorm:"size:100;not null" json:"name"`
	Description  string  `gorm:"type:text" json:"description"`
	Price        float64 `gorm:"not null" json:"price"`
	Category     string  `gorm:"size:50;not null" json:"category"`
	Availability bool    `gorm:"default:true" json:"availability"`
	Tags         string  `gorm:"type:text" json:"tags"`
}
