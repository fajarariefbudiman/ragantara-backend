package handlers

import (
	"fmt"
	"net/http"
	"real-time-application/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
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
	result, err := services.GetProducts(pageInt, limitInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": err,
		})
	}
	return c.JSON(http.StatusOK, result)
}

func GetProductBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := services.GetProductBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": err,
		})
	}
	return c.JSON(http.StatusOK, result)
}

func CreateProduct(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	body := c.FormValue("body")
	price := c.FormValue("price")
	convertPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Error Parse float 64")
		return err
	}
	slug := c.FormValue("slug")
	category_id := c.FormValue("category_id")
	convertCategoryId, err := strconv.Atoi(category_id)
	if err != nil {
		fmt.Println("Error Parse int Category_Id")
		return err
	}
	quantity := c.FormValue("quantity")
	convertQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		return err
	}
	discount := c.FormValue("discount")
	convertDiscount, err := strconv.Atoi(discount)
	if err != nil {
		return err
	}

	result, err := services.CreateProduct(name, description, body, slug, convertPrice, convertQuantity, convertCategoryId, convertDiscount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": err,
		})
	}
	return c.JSON(http.StatusCreated, result)
}

func UpdateProduct(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	body := c.FormValue("body")
	price := c.FormValue("price")
	convertPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Error Parse float 64")
		return err
	}
	slug := c.FormValue("slug")
	category_id := c.FormValue("category_id")
	convertCategoryId, err := strconv.Atoi(category_id)
	if err != nil {
		fmt.Println("Error Parse int Category_Id")
		return err
	}
	quantity := c.FormValue("quantity")
	convertQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		return err
	}
	discount := c.FormValue("discount")
	convertDiscount, err := strconv.Atoi(discount)
	if err != nil {
		return err
	}

	result, err := services.UpdateProduct(name, description, body, slug, convertPrice, convertQuantity, convertCategoryId, convertDiscount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": err,
		})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteProduct(c echo.Context) error {
	id := c.FormValue("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error Parse int Id")
		return err
	}

	result, err := services.DeleteProduct(convertId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Error": err,
		})
	}
	return c.JSON(http.StatusNoContent, result)
}
