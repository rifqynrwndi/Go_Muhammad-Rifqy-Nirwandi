package config

import (
    "project/models"
)

func MigrateDB() {
    DB.AutoMigrate(&models.User{})
    DB.AutoMigrate(&models.Package{})
}
