package handlers

import (
	"net/http"
	"real-time-application/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllCourts(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	pageInt, _ := strconv.Atoi(page)
	if pageInt <= 0 {
		pageInt = 1
	}
	limitInt, _ := strconv.Atoi(limit)
	if limitInt <= 0 {
		limitInt = 7
	}
	result, err := services.GetAllCourts(pageInt, limitInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"disini":  "disini",
		})
	}
	return c.JSON(http.StatusOK, result)
}

func GetCourtById(c echo.Context) error {
	id := c.Param("id")
	str_id, _ := strconv.Atoi(id)
	result, err := services.GetCourtById(str_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": err,
		})
	}
	return c.JSON(http.StatusOK, result)
}
