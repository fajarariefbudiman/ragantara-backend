package services

import (
	"net/http"
	"real-time-application/database"
	"time"
)

type Cart struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	User_Id    int       `json:"user_id" gorm:"not null"`
	Product_Id int       `json:"product_id" gorm:"not null"`
	Quantity   int       `json:"quantity" gorm:"default:1"`
	Created_At time.Time `json:"created_at" gorm:"autoCreateTime;type:timestamp"`
	Updated_At time.Time `json:"updated_at" gorm:"autoUpdateTime;type:timestamp"`
	User       User      `json:"user" gorm:"foreignKey:User_Id;constraint:OnDelete:CASCADE"`
	Product    Product   `json:"product" gorm:"foreignKey:Product_Id;constraint:OnDelete:CASCADE"`
}

func GetCartItems(user_id int) (Response, error) {
	var res Response
	var cartItems []Cart
	db := database.DB
	if err := db.Preload("Product").Preload("User").Where("user_id = ?", user_id).Find(&cartItems).Error; err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Get all cart items"
	res.Data = cartItems
	return res, nil
}

func AddToCart(user_id, product_id, quantity int) (Response, error) {
	var res Response
	db := database.DB
	cart := Cart{
		User_Id:    user_id,
		Product_Id: product_id,
		Quantity:   quantity,
	}
	result := db.Create(&cart)
	if result.Error != nil {
		return res, result.Error
	}
	res.Status = http.StatusCreated
	res.Message = "Item added to cart"
	res.Data = cart
	return res, nil
}

func UpdateCartItem(user_id, product_id, quantity int) (Response, error) {
	var res Response
	db := database.DB
	cart := Cart{}
	if err := db.Where("user_id = ? AND product_id = ?", user_id, product_id).First(&cart).Error; err != nil {
		return res, err
	}
	cart.Quantity = quantity
	if err := db.Save(&cart).Error; err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Cart item updated"
	res.Data = cart
	return res, nil
}

func DeleteCartItem(user_id, product_id int) (Response, error) {
	var res Response
	db := database.DB
	result := db.Where("user_id = ? AND product_id = ?", user_id, product_id).Delete(&Cart{})
	if result.Error != nil {
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "Cart item removed"
	res.Data = result.RowsAffected
	return res, nil
}
