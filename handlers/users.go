package handlers

import (
	"net/http"
	"projet-go/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserHandler struct
type UserHandler struct {
	DB *gorm.DB
}

// CreateUser creates a new user
func CreateUser(c echo.Context) error {
	// Bind request body to user struct
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	// Hash user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Create user
	result := models.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	// Remove password from response
	user.Password = ""

	return c.JSON(http.StatusCreated, user)
}

// GetUsers returns all users
func GetUsers(c echo.Context) error {
	// Get all users
	users := []models.User{}
	result := models.DB.Preload("Roles").Preload("Groups").Find(&users)
	if result.Error != nil {
		return result.Error
	}

	// Remove password from response
	for i := range users {
		users[i].Password = ""
	}

	return c.JSON(http.StatusOK, users)
}

// UpdateUser updates an existing user
func UpdateUser(c echo.Context) error {
	// Get user ID from URL parameter
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Get user from database
	user := models.User{}
	result := models.DB.Preload("Roles").Preload("Groups").First(&user, userID)
	if result.Error != nil {
		return result.Error
	}

	// Bind request body to user struct
	updatedUser := models.User{}
	if err := c.Bind(&updatedUser); err != nil {
		return err
	}

	// Update user fields
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Roles = updatedUser.Roles
	user.Groups = updatedUser.Groups

	// Save changes to database
	result = models.DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}

	// Remove password from response
	user.Password = ""

	return c.JSON(http.StatusOK, user)
}

// DeleteUser deletes an existing user
func DeleteUser(c echo.Context) error {
	// Get user ID from URL parameter
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Delete user from database
	result := models.DB.Delete(&models.User{}, userID)
	if result.Error != nil {
		return result.Error
	}

	return c.NoContent(http.StatusNoContent)
}

// Create an instance of the Echo framework

// Initialize routes
func UsersHandler(e *echo.Echo) {

	e.GET("", GetRoles)
	e.POST("", CreateUser)
	e.PUT("/:id", UpdateRole)
	e.DELETE("/:id", DeleteRole)
}
