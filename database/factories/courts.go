package factories

import (
	"fmt"
	"math"
	"math/rand"
	"real-time-application/database"
	"real-time-application/models"

	"github.com/bxcodec/faker/v4"
)

func CourtFactory() *models.Court {
	courtType := []string{"Indoor", "Outdoor"}
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
	numFacilities := rand.Intn(8) + 1

	selectedFacilities := make([]models.Facility, numFacilities)
	for i := 0; i < numFacilities; i++ {
		selectedFacilities[i] = facilities[rand.Intn(len(facilities))]
	}

	openingHour := rand.Intn(8) + 6
	closingHour := openingHour + rand.Intn(6)
	return &models.Court{
		Name:           faker.Word(),
		Location_Id:    rand.Intn(8) + 1,
		Price_Per_Hour: rand.Float64()*50000 + 100000,
		Court_Type:     courtType[rand.Intn(len(courtType))],
		Opening_Hours:  fmt.Sprintf("%01d:00", openingHour),
		Closing_Hours:  fmt.Sprintf("%01d:00", closingHour),
		Rating:         math.Round((rand.Float64()*4+1)*100) / 100,
		Review_Count:   rand.Intn(100),
		ImageURL:       faker.URL(),
		Facilities:     selectedFacilities,
	}
}

func CreateDataCourts(count int) {
	for i := 0; i < count; i++ {
		court := CourtFactory()
		if err := database.DB.Create(&court).Error; err != nil {
			panic(err)
		}
	}
	fmt.Println("Success Create Factory Courts")
}
