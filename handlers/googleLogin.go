package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"real-time-application/config"
	"real-time-application/database"
	"real-time-application/services"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func SplitName(name string) (string, string) {
	parts := strings.Split(name, " ")
	firstname := parts[0]
	lastname := ""

	if len(parts) > 1 {
		lastname = parts[1]
	}

	return firstname, lastname

}

func GenerateState() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func HandleGoogleLogin(c echo.Context) error {
	// fmt.Println("Google Login Handler hit!")
	state := GenerateState()
	url := config.GoogleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	// fmt.Println("Generated URL:", url)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	// log.Println("Code received:", code)
	if code == "" {
		return c.String(http.StatusBadRequest, "Missing code parameter")
	}

	token, err := config.GoogleOauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to exchange token")
	}

	client := config.GoogleOauthConfig.Client(c.Request().Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get user info")
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to decode user info")
	}

	email := userInfo["email"].(string)
	name := userInfo["name"].(string)

	existingUser := services.User{}
	result := database.DB.Where("email = ?", email).First(&existingUser)

	var user services.User
	if result.Error == nil {
		user = existingUser
	} else {
		firstname, lastname := SplitName(name)
		newUser := services.User{
			Firstname: firstname,
			Lastname:  lastname,
			Email:     email,
		}

		if err := database.DB.Create(&newUser).Error; err != nil {
			return c.String(http.StatusInternalServerError, "Failed to create user")
		}
		user = newUser
	}
	// fmt.Println(user)
	authToken := jwt.New(jwt.SigningMethodHS256)
	claims := authToken.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := authToken.SignedString(SigningKey)
	fmt.Println(t)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed Generate Token",
			"error":   err.Error(),
		})
	}

	c.SetCookie(&http.Cookie{
		Name:     "authToken",
		Value:    t,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})
	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")

}

func HandleFacebookLogin(c echo.Context) error {
	url := config.FacebookOauthConfig.AuthCodeURL("state")
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleFacebookCallback(c echo.Context) error {
	code := c.QueryParam("code")
	token, err := config.FacebookOauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to exchange token")
	}

	client := config.FacebookOauthConfig.Client(c.Request().Context(), token)
	resp, err := client.Get("https://graph.facebook.com/me?fields=id,name,email")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get user info")
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to decode user info")
	}

	// Di sini Anda bisa menambahkan logic untuk:
	// 1. Menyimpan user ke database
	// 2. Generate JWT token
	// 3. Menyimpan session

	redirectURL := fmt.Sprintf("http://localhost:3000/auth/callback?user=%s&email=%s",
		userInfo["name"], userInfo["email"])
	return c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
