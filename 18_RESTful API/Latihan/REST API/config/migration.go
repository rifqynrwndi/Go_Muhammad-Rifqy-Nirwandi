package config

import "rest/models"

func MigrateDB() {
	DB.AutoMigrate(&models.Books{})
}