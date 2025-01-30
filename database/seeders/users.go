package seeders

import (
	"fmt"
	"real-time-application/database"
	"real-time-application/services"

	"golang.org/x/crypto/bcrypt"
)

func UserSeeder() {
	password, err := bcrypt.GenerateFromPassword([]byte("123123123"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	users := []services.User{
		{Firstname: "Fajar", Lastname: "Arief", Age: 19, Address: "Tangerang", Email: "budimanfajar660@gmail.com", Password: string(password)},
		{Firstname: "admin1", Lastname: "__", Age: 19, Address: "Tangerang", Email: "budimanfajar661@gmail.com", Password: string(password), Role: "admin"},
		{Firstname: "Fajar", Lastname: "Arief Budiman", Age: 19, Address: "Tangerang", Email: "budimanfajar662@gmail.com", Password: string(password), Role: "player"},
	}

	for _, user := range users {
		if err := database.DB.Create(&user).Error; err != nil {
			fmt.Println("Create users")
		}
	}

	fmt.Println("Success Create users")
}
