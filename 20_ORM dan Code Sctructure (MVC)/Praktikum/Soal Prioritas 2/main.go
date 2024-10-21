package main

import (
	"billing/config"
	"billing/routes"
)

func main() {
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	r := routes.SetupRoutes(db)
	r.Logger.Fatal(r.Start(":8000"))
}
