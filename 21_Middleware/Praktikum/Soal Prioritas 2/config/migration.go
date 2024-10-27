package config

import (
    "posts/models"
)

func MigrateDB() {
    DB.AutoMigrate(&models.User{})
    DB.AutoMigrate(&models.Post{})
}
