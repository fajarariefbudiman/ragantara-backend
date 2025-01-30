package services

import (
	"net/http"
	"real-time-application/database"
	"time"
)

type Category struct {
	Id          int       `validate:"required" json:"id"`
	Name        string    `validate:"required" json:"name"`
	Slug        string    `validate:"required" json:"slug"`
	Description string    `validate:"required" json:"description"`
	Created_At  time.Time `json:"created_at" gorm:"autoCreateTime;type:timestamp"`
	Updated_At  time.Time `json:"updated_at" gorm:"autoUpdateTime;type:timestamp"`
}

func GetCategories() (Response, error) {
	var res Response
	categories := []Category{}
	db := database.DB
	sql := []string{"name", "slug", "description"}
	if err := db.Select(sql).Find(&categories).Error; err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "get all categories"
	res.Data = categories
	return res, nil
}

func CreateCategory(name, description, slug string) (Response, error) {
	var response Response
	db := database.DB
	// fmt.Println("Success Connect To Database")
	category := &Category{
		Name:        name,
		Description: description,
		Slug:        slug,
	}
	result := db.Create(&category)

	response.Message = "success create"
	response.Status = 201
	response.Data = result.RowsAffected

	return response, nil
}
