package seeders

import (
	"fmt"
	"real-time-application/database"
	"real-time-application/models"
)

func LocationSeeder() {
	locations := []models.Location{
		{Id: 1, Name: "Kab.Tangerang"},
		{Id: 2, Name: "Tangerang Selatan"},
		{Id: 3, Name: "Tangerang Kota"},
		{Id: 4, Name: "Jakarta Selatan"},
		{Id: 5, Name: "Jakarta Barat"},
		{Id: 6, Name: "Jakarta Pusat"},
		{Id: 7, Name: "Jakarta Utara"},
		{Id: 8, Name: "Jakarta Timur"},
	}
	for _, location := range locations {
		if err := database.DB.Create(&location).Error; err != nil {
			fmt.Println("Create Locations Error")
		}
	}
}
