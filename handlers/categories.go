package handlers

import (
	"net/http"
	"real-time-application/services"

	"github.com/labstack/echo/v4"
)

func GetCategories(c echo.Context) error {
	result, err := services.GetCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": err,
		})
	}
	return c.JSON(http.StatusOK, result)
}

func CreateCategory(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	slug := c.FormValue("slug")
	result, err := services.CreateCategory(name, description, slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": err,
		})
	}
	return c.JSON(http.StatusCreated, result)
}
