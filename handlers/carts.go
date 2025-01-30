package handlers

import (
	"net/http"
	"real-time-application/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddToCart(c echo.Context) error {
	userID := c.FormValue("user_id")
	productID := c.FormValue("product_id")
	quantity := c.FormValue("quantity")

	convertUserID, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid user_id",
		})
	}

	convertProductID, err := strconv.Atoi(productID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid product_id",
		})
	}

	convertQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid quantity",
		})
	}

	result, err := services.AddToCart(convertUserID, convertProductID, convertQuantity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to add to cart",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, result)
}

func GetCartItems(c echo.Context) error {
	userID := c.QueryParam("user_id")
	convertUserID, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid user_id",
		})
	}

	items, err := services.GetCartItems(convertUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to retrieve cart items",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, items)
}

func DeleteCartItem(c echo.Context) error {
	userID := c.Param("user_id")
	productID := c.Param("product_id")

	convertUserID, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid user_id",
		})
	}

	convertProductID, err := strconv.Atoi(productID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid product_id",
		})
	}

	result, err := services.DeleteCartItem(convertUserID, convertProductID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete cart item",
			"error":   err.Error(),
		})
	}
	if result.Data == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "No matching cart item found to delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Cart item deleted successfully",
	})
}
