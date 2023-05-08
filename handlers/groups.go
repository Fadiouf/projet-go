package handlers

import (
	"net/http"
	"strconv"

	"projet-go/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// GroupHandler struct
type GroupHandler struct {
	DB *gorm.DB
}

// CreateGroup creates a new group
func CreateGroup(c echo.Context) error {
	// Bind request body to group struct
	group := models.Group{}
	if err := c.Bind(&group); err != nil {
		return err
	}

	// Create group
	result := models.DB.Create(&group)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusCreated, group)
}

// GetUsers returns all groups
func GetGroups(c echo.Context) error {
	// Get all groups
	groups := []models.Group{}
	result := models.DB.Preload("Roles").Preload("Users").Find(&groups)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, groups)
}

// UpdateGroup updates an existing group
func UpdateGroup(c echo.Context) error {
	// Get group ID from URL parameter
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Get group from database
	group := models.Group{}
	result := models.DB.Preload("Roles").Preload("Users").First(&group, groupID)
	if result.Error != nil {
		return result.Error
	}

	// Bind request body to group struct
	updatedGroup := models.Group{}
	if err := c.Bind(&updatedGroup); err != nil {
		return err
	}

	// Update group fields
	group.Name = updatedGroup.Name
	group.Roles = updatedGroup.Roles
	group.Users = updatedGroup.Users

	// Save changes to database
	result = models.DB.Save(&group)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, group)
}

// DeleteGroup deletes an existing group
func DeleteGroup(c echo.Context) error {
	// Get group ID from URL parameter
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Delete group from database
	result := models.DB.Delete(&models.Group{}, groupID)
	if result.Error != nil {
		return result.Error
	}

	return c.NoContent(http.StatusNoContent)
}

// Initialize routes
func GroupsHandler(e *echo.Echo) {

	e.GET("", GetGroups)
	e.POST("", CreateGroup)
	e.PUT("/:id", UpdateGroup)
	e.DELETE("/:id", DeleteGroup)
}
