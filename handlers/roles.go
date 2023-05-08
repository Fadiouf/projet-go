package handlers

import (
	"net/http"
	"strconv"

	"projet-go/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RoleHandler struct
type RoleHandler struct {
	DB *gorm.DB
}

// CreateRole creates a new role
func CreateRole(c echo.Context) error {
	// Bind request body to role struct
	role := models.Role{}
	if err := c.Bind(&role); err != nil {
		return err
	}

	// Create role
	result := models.DB.Create(&role)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusCreated, role)
}

// GetRoles returns all roles
func GetRoles(c echo.Context) error {
	// Get all roles
	roles := []models.Role{}
	result := models.DB.Preload("Users").Preload("Groups").Find(&roles)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, roles)
}

// UpdateRole updates an existing role
func UpdateRole(c echo.Context) error {
	// Get role ID from URL parameter
	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Get role from database
	role := models.Role{}
	result := models.DB.Preload("Users").Preload("Groups").First(&role, roleID)
	if result.Error != nil {
		return result.Error
	}

	// Bind request body to role struct
	updatedRole := models.Role{}
	if err := c.Bind(&updatedRole); err != nil {
		return err
	}

	// Update user fields
	role.Name = updatedRole.Name
	role.Users = updatedRole.Users
	role.Groups = updatedRole.Groups
	role.Description = updatedRole.Description

	// Save changes to database
	result = models.DB.Save(&role)
	if result.Error != nil {
		return result.Error
	}

	return c.JSON(http.StatusOK, role)
}

// DeleteRoles deletes an existing role
func DeleteRole(c echo.Context) error {
	// Get role ID from URL parameter
	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// Delete role from database
	result := models.DB.Delete(&models.Role{}, roleID)
	if result.Error != nil {
		return result.Error
	}

	return c.NoContent(http.StatusNoContent)
}

// Initialize routes
func RolesHandler(e *echo.Echo) {

	e.GET("", GetUsers)
	e.POST("", CreateRole)
	e.PUT("/:id", UpdateUser)
	e.DELETE("/:id", DeleteUser)
}
