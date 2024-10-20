package models

type Books struct {
	Id     int    `json:"id" form:"id" gorm:"primaryKey"`
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
}
