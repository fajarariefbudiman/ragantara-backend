package handlers

import (
	"fmt"
	"net/http"
	"real-time-application/services"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Email string `json:"email"`
	*jwt.StandardClaims
}

var SigningKey = []byte("secret")

func Login(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}
	// fmt.Println("Received email:", req.Email)
	// fmt.Println("Password hash:", req.Password)

	user, err := services.AuthUser(req.Email, req.Password)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
			"error":   err.Error(),
		})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	t, err := token.SignedString(SigningKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed Generate Token",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login Successful",
		"token":   t,
		"exp":     claims["exp"],
		"auth":    user,
	})
}

func Register(c echo.Context) error {
	var req struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error generating password",
			"error":   err.Error(),
		})
	}

	result, err := services.Register(req.Firstname, req.Lastname, req.Email, string(hashPassword))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"error":   err.Error(),
		})
	}
	// fmt.Printf("Data yang dikirim", result)
	return c.JSON(http.StatusCreated, result)
}
