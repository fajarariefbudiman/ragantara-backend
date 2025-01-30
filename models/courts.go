package models

import "time"

type Court struct {
	Id             int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string     `gorm:"type:varchar(255);not null" json:"name"`
	Location_Id    int        `gorm:"not null" json:"location_id"`
	Price_Per_Hour float64    `gorm:"type:decimal(10,2);not null" json:"price_per_hour"`
	Court_Type     string     `gorm:"type:enum('Indoor','Outdoor');not null" json:"court_type"`
	Opening_Hours  string     `gorm:"type:time;not null" json:"opening_hours"`
	Closing_Hours  string     `gorm:"type:time;not null" json:"closing_hours"`
	Rating         float64    `gorm:"type:decimal(3,2);default:0.00" json:"rating"`
	Review_Count   int        `gorm:"default:0" json:"review_count"`
	ImageURL       string     `gorm:"type:varchar(255);not null" json:"image_url"`
	Facilities     []Facility `gorm:"many2many:court_facilities;" json:"facilities"`
	Created_At     time.Time  `gorm:"autoCreateTime" json:"created_at"`
	Updated_At     time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
