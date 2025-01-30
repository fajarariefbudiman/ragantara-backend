package seeders

import (
	"fmt"
	"real-time-application/database"
	"real-time-application/models"
)

func FacilitySeeder() {
	facilities := []models.Facility{
		{Id: 1, Name: "Ruang Ganti"},
		{Id: 2, Name: "Area Parkir"},
		{Id: 3, Name: "Air Mineral"},
		{Id: 4, Name: "Pencahayaan"},
		{Id: 5, Name: "Cafe"},
		{Id: 6, Name: "Kamar Mandi"},
		{Id: 7, Name: "Mushola"},
		{Id: 8, Name: "CCTV"},
		{Id: 9, Name: "P3K"},
	}
	for _, facility := range facilities {
		if err := database.DB.Create(&facility).Error; err != nil {
			fmt.Println("Create Facilities")
		}
	}
}
