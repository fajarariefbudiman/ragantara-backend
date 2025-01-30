package services

import (
	"fmt"
	"net/http"
	"real-time-application/database"
	"time"
)

type Product struct {
	Id          int       `json:"id"`
	Name        string    `validate:"required" json:"name"`
	Description string    `validate:"required" json:"description"`
	Body        string    `validate:"required" json:"body"`
	Slug        string    `validate:"required" json:"slug"`
	Price       float64   `validate:"required" json:"price"`
	Quantity    int       `validate:"required" json:"quantity"`
	Category_Id int       `validate:"required" json:"category_id"`
	Discount    int       `validate:"required" json:"discount"`
	Created_At  time.Time `json:"created_at" gorm:"autoCreateTime;type:timestamp"`
	Updated_At  time.Time `json:"updated_at" gorm:"autoUpdateTime;type:timestamp"`
	Category    Category  `json:"category" gorm:"foreignKey:Category_Id;constraint:OnDelete:CASCADE"`
}

func GetProducts(page, limit int) (Response, error) {
	var res Response
	offset := (page - 1) * limit
	products := []Product{}
	db := database.DB
	if err := db.Preload("Category").Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return res, err
	}

	var totalItems int64
	if err := db.Model(&Product{}).Count(&totalItems).Error; err != nil {
		return res, err
	}

	totalPage := int(totalItems) / limit
	if int(totalItems)%limit != 0 {
		totalPage++
	}

	res.Status = http.StatusOK
	res.Message = "get all products"
	res.Data = map[string]interface{}{
		"products":    products,
		"totalItems":  totalItems,
		"totalPage":   totalPage,
		"currentPage": page,
	}
	return res, nil
}

func GetProductBySlug(slug string) (Response, error) {
	var res Response
	product := Product{}
	db := database.DB
	column := []string{"name", "description", "body", "slug"}
	if err := db.Preload("Category").Select(column).Where("slug = ?", slug).First(&product).Error; err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "get product"
	res.Data = product
	return res, nil
}

func CreateProduct(name, description, body, slug string, price float64, quantity, category_id, discount int) (Response, error) {
	var response Response
	db := database.DB
	fmt.Println("Success Connect To Database")
	products := &Product{
		Name:        name,
		Description: description,
		Body:        body,
		Slug:        slug,
		Price:       price,
		Quantity:    quantity,
		Category_Id: category_id,
		Discount:    discount,
	}
	result := db.Create(&products)

	response.Message = "Success Create"
	response.Status = 201
	response.Data = result.RowsAffected

	return response, nil
}

func UpdateProduct(name, description, body, slug string, price float64, quantity, category_id, discount int) (Response, error) {
	var response Response
	db := database.DB
	fmt.Println("Success Connect To Database")
	products := &Product{
		Name:        name,
		Description: description,
		Body:        body,
		Slug:        slug,
		Price:       price,
		Quantity:    quantity,
		Category_Id: category_id,
		Discount:    discount,
	}
	result := db.Model(&Product{}).Where("slug = ?", slug).Save(products)

	response.Message = "Success Create"
	response.Status = 201
	response.Data = result.RowsAffected

	return response, nil
}

func DeleteProduct(id int) (Response, error) {
	var response Response
	db := database.DB
	result := db.Model(&Product{}).Where("id = ?", id).Delete(&Product{})

	response.Message = "Success Create"
	response.Status = 201
	response.Data = result.RowsAffected

	return response, nil
}
