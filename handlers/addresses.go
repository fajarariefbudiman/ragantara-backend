package handlers

import (
	"fmt"
	"net/http"
	"real-time-application/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAddresses(c echo.Context) error {
	user := c.QueryParam("user_id")
	fmt.Println("Received user_id:", user)

	user_id, err := strconv.Atoi(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"error":   err.Error(),
		})
	}
	result, err := services.GetAddresses(user_id)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
			"error":   err.Error(),
		})
	}
	fmt.Println("Result", result)
	return c.JSON(http.StatusOK, result)
}

func CreateAddress(c echo.Context) error {
	var req struct {
		User_Id        int    `json:"user_id"`
		Address_Type   string `json:"address_type" `
		Recipient_Name string `json:"recipient_name"`
		Phone_Number   string `json:"phone_number"`
		Full_Address   string `json:"full_address"`
		District       string `json:"district"`
		City           string `json:"city"`
		Province       string `json:"province"`
		Postal_Code    string `json:"postal_code"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error generating password",
			"error":   err.Error(),
		})
	}

	result, err := services.CreateAddress(req.User_Id, req.Address_Type, req.Recipient_Name, req.Phone_Number, req.Full_Address, req.District, req.City, req.Province, req.Postal_Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error Fail",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, result)
}
