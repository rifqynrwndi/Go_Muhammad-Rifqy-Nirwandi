package config

import (
	"postcategory/models"
)

func MigrateDB() {
    DB.AutoMigrate(&models.User{})
    DB.AutoMigrate(&models.Post{})
    DB.AutoMigrate(&models.Category{})
}
