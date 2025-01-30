package routes

import (
	"net/http"
	"real-time-application/handlers"
	"real-time-application/middleware"

	"github.com/labstack/echo/v4"
)

func API() *echo.Echo {
	e := echo.New()
	middleware.AuthMiddleware(e)
	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.POST("/api/auth/login", handlers.Login)
	e.POST("/api/auth/register", handlers.Register)
	e.GET("/api/products", handlers.GetProducts)
	e.GET("/api/categories", handlers.GetCategories)
	e.GET("/api/courts", handlers.GetAllCourts)

	e.GET("/auth/google", handlers.HandleGoogleLogin)
	e.GET("/auth/google/callback", handlers.HandleGoogleCallback)
	e.GET("/auth/facebook", handlers.HandleFacebookLogin)
	e.GET("/auth/facebook/callback", handlers.HandleFacebookCallback)

	// Protected routes
	api := e.Group("/api")
	api.Use(middleware.JWTMiddleware())

	// Products routes (protected)
	products := api.Group("/product")
	products.GET("/:slug", handlers.GetProductBySlug)
	products.POST("", handlers.CreateProduct)
	products.PUT("/:slug", handlers.UpdateProduct)
	products.DELETE("/:id", handlers.DeleteProduct)

	// Courts
	courts := api.Group("/courts")
	courts.GET("/:id", handlers.GetCourtById)

	// Categories routes (protected)
	categories := api.Group("/categories")
	categories.POST("", handlers.CreateCategory)

	// addresses
	addresses := api.Group("/addresses")
	addresses.GET("", handlers.GetAddresses)
	addresses.POST("", handlers.CreateAddress)

	// Orders

	// Checkout

	return e
}
