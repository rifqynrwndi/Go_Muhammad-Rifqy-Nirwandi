package models

type Post struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
}
