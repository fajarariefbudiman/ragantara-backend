package main

import (
	"log"
	"real-time-application/config"
	"real-time-application/database"
	"real-time-application/routes"

	"github.com/joho/godotenv"
)

func main() {
	database.ConfigDatabase()
	config.Init()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// seeders.CategorySeeder()
	// seeders.UserSeeder()
	// seeders.LocationSeeder()
	// seeders.FacilitySeeder()
	// seeders.AddressesSeeder()
	// factories.CreateDataProducts(30)
	// factories.CreateDataCourts(30)

	api := routes.API()

	api.Logger.Fatal(api.Start(":1323"))
}
