package factories

import (
	"fmt"
	"math/rand"
	"real-time-application/database"
	"real-time-application/services"

	"github.com/bxcodec/faker/v4"
)

func ProductFactory() *services.Product {
	return &services.Product{
		Name:        faker.Word(),
		Description: faker.Sentence(),
		Body:        faker.Paragraph(),
		Slug:        faker.Word(),
		Price:       rand.Float64() + 100000,
		Quantity:    rand.Intn(10) + 1,
		Category_Id: rand.Intn(10) + 1,
		Discount:    rand.Intn(50) + 1,
	}
}

func CreateDataProducts(count int) {
	for i := 0; i < count; i++ {
		product := ProductFactory()
		if err := database.DB.Create(&product).Error; err != nil {
			panic(err)
		}
	}
	fmt.Println("Success Create Factory Products")
}
