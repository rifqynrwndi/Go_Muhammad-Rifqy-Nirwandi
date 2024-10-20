package main

import (
	"database/sql"
	"log"

	"billing/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func initDB() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/billing_db")
    if err != nil {
        return nil, err
    }
    return db, nil
}

func main() {
    db, err := initDB()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    e := echo.New()
    routes.SetupRoutes(e, db)

    log.Fatal(e.Start(":8000"))
}
