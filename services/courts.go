package services

import (
	"net/http"
	"real-time-application/database"
	"real-time-application/models"
)

func GetAllCourts(page, limit int) (Response, error) {
	var res Response
	db := database.DB
	courts := []models.Court{}
	facilities := []models.Facility{}
	locations := []models.Location{}
	offset := (page - 1) * limit
	if err := db.Model(&models.Facility{}).Find(&facilities).Error; err != nil {
		return res, err
	}
	if err := db.Model(&models.Location{}).Find(&locations).Error; err != nil {
		return res, err
	}
	if err := db.Preload("Facilities").Offset(offset).Limit(limit).Find(&courts).Error; err != nil {
		return res, err
	}
	var totalItems int64
	if err := db.Model(&models.Court{}).Count(&totalItems).Error; err != nil {
		return res, err
	}

	totalPage := totalItems / int64(limit)
	if totalItems%int64(limit) != 0 {
		totalPage++
	}
	res.Status = http.StatusOK
	res.Message = "Success Get All Courts"
	res.Data = map[string]interface{}{
		"totalItems":  totalItems,
		"totalPage":   totalPage,
		"locations":   locations,
		"facilities":  facilities,
		"courts":      courts,
		"currentPage": page,
	}
	return res, nil

}

func GetCourtById(id int) (Response, error) {
	var res Response
	court := models.Court{}
	db := database.DB
	if err := db.Preload("Facilities").Select("*").Where("id = ?", id).First(&court).Error; err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "get court"
	res.Data = court
	return res, nil
}
