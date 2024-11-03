package config

import (
    "fmt"
    "go-wishlist-api/entities"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "os"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDatabase() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("Error loading .env file")
    }

    var err error

    dbUser := os.Getenv("DATABASE_USER")
    dbPassword := os.Getenv("DATABASE_PASSWORD")
    dbHost := os.Getenv("DATABASE_HOST")
    dbPort := os.Getenv("DATABASE_PORT")
    dbName := os.Getenv("DATABASE_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPassword, dbHost, dbPort, dbName,
    )

    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    DB.AutoMigrate(&entities.Wishlist{})
}
