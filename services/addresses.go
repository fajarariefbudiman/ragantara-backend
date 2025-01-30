package services

import (
	"net/http"
	"real-time-application/database"
	"time"
)

type Address struct {
	Id             int       `json:"id"`
	User_Id        int       `json:"user_id" gorm:"not null"`
	Address_Type   string    `json:"address_type" gorm:"type:enum('Rumah', 'Kantor');not null"`
	Recipient_Name string    `json:"recipient_name" gorm:"type:varchar(255);not null"`
	Phone_Number   string    `json:"phone_number" gorm:"type:varchar(25);not null"`
	Full_Address   string    `json:"full_address" gorm:"type:text;not null"`
	District       string    `json:"district" gorm:"type:varchar(100);not null"`
	City           string    `json:"city" gorm:"type:varchar(100);not null"`
	Province       string    `json:"province" gorm:"type:varchar(100);not null"`
	Postal_Code    string    `json:"postal_code" gorm:"type:varchar(20);not null"`
	Is_Primary     bool      `json:"is_primary" gorm:"default:false"`
	Created_At     time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_At     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func GetAddresses(user_id int) (Response, error) {
	var res Response
	addresses := []Address{}
	db := database.DB
	column := []string{"address_type", "recipient_name", "city", "phone_number", "postal_code", "full_address", "district", "province", "is_primary"}
	if err := db.Select(column).Where("user_id = ?", user_id).Find(&addresses).Error; err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "success get addresses user"
	res.Data = addresses

	return res, nil
}

func CreateAddress(user_id int, address_type, recipient_name, phone_number, full_address, district, city, province, postal_code string) (Response, error) {
	var res Response
	db := database.DB
	address := Address{
		User_Id:        user_id,
		Address_Type:   address_type,
		Recipient_Name: recipient_name,
		Phone_Number:   phone_number,
		Full_Address:   full_address,
		District:       district,
		City:           city,
		Province:       province,
		Postal_Code:    postal_code,
	}
	result := db.Create(&address)

	res.Status = http.StatusCreated
	res.Message = "success create"
	res.Data = result.RowsAffected

	return res, nil
}
