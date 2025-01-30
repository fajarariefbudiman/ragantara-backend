package seeders

import (
	"fmt"
	"real-time-application/database"
	"real-time-application/models"
)

func CourtsSeeder() {
	courts := []models.Court{
		{Name: "Bintang Futsal", Location_Id: 1, Price_Per_Hour: 178.0000, Court_Type: "indoor"},
		{Name: "King Futsal", Location_Id: 2, Price_Per_Hour: 128.0000, Court_Type: "indoor"},
		{Name: "Imperial Futsal", Location_Id: 3, Price_Per_Hour: 148.0000, Court_Type: "indoor"},
		{Name: "Vero Futsal", Location_Id: 4, Price_Per_Hour: 181.0000, Court_Type: "indoor"},
		{Name: "Temon", Location_Id: 5, Price_Per_Hour: 120.0000, Court_Type: "indoor"},
		{Name: "Asetsa Futsal", Location_Id: 6, Price_Per_Hour: 133.0000, Court_Type: "indoor"},
		{Name: "Aviat", Location_Id: 7, Price_Per_Hour: 210.0000, Court_Type: "indoor"},
		{Name: "Vegas Futsal", Location_Id: 8, Price_Per_Hour: 118.0000, Court_Type: "indoor"},
		{Name: "Prima", Location_Id: 1, Price_Per_Hour: 138.0000, Court_Type: "indoor"},
		{Name: "Dream Futsal", Location_Id: 2, Price_Per_Hour: 140.0000, Court_Type: "indoor"},
		{Name: "Futsal Champ", Location_Id: 3, Price_Per_Hour: 150.0000, Court_Type: "indoor"},
	}
	for _, court := range courts {
		if err := database.DB.Create(court).Error; err != nil {
			fmt.Println("Error")
		}
	}
}
