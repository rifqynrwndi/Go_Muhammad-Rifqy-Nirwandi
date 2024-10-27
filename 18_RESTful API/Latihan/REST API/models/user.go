package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}