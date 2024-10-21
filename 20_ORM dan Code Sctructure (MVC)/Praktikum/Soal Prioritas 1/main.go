package main

import (
	"app/config"
	"app/routes"
)

func main() {
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	r := routes.SetupRoutes(db)
	r.Logger.Fatal(r.Start(":8080"))
}
