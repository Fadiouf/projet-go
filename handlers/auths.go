package handlers

import (
	"net/http"
	"projet-go/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// AuthHandler struct
type AuthHandler struct {
	DB *gorm.DB
}

// Authenticate creates a new auth
func Authenticate(c echo.Context) error {
	// Bind request body to auth struct
	auth := models.AuthToken{}
	if err := c.Bind(&auth); err != nil {
		return err
	}

	// Create auth
	result := models.DB.Create(&auth)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusCreated, auth)
}

// Initialize routes
func AuthsHandler(e *echo.Echo) {

	e.POST("", Authenticate)
}
