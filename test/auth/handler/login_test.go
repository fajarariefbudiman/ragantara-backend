package handler_test

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"real-time-application/handlers"
// 	"real-time-application/services"
// 	"strings"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )

// type MockAuth struct{}

// func (m *MockAuth) Auth(email, password string) (*services.User, error) {
// 	if email == "budimanfajar660@gmail.com" && password == "password" {
// 		return &services.User{
// 			Email:    email,
// 			Password: password,
// 		}, nil
// 	}
// 	return nil, fmt.Errorf("Invalid credentials")
// }

// func TestLoginHandler(t *testing.T) {
// 	e := echo.New()
// 	loginRequest := map[string]string{
// 		"email":    "budimanfajar660@gmail.com",
// 		"password": "password",
// 	}
// 	reqBody, _ := json.Marshal(loginRequest)
// 	req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(string(reqBody)))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	mockService := &MockAuth{}
// 	handler := handlers.NewHandler(mockService)

// 	err := handler.Login(c)
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		var response map[string]interface{}
// 		json.Unmarshal(rec.Body.Bytes(), &response)
// 		assert.Equal(t, "Login Successful", response["message"])
// 		assert.NotEmpty(t, response["token"])
// 	}
// }
