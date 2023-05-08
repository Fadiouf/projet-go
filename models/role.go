package models

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// Role model
type Role struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"uniqueIndex;not null"`
	Description string    `json:"description" gorm:"not null"`
	Created_At  time.Time `json:"created" gorm:"not null"`
	Updated_At  time.Time `json:"updated_at" gorm:"not null"`
	Deleted_At  time.Time `json:"deleted_at" gorm:"index"`
	Users       []*User   `json:"users" gorm:"many2many:user_roles;"`
	Groups      []*Group  `json:"groups" gorm:"many2many:group_roles;"`
}

// Handle get roles
func handleGetRoles(c echo.Context) error {
	var roles []Role
	if err := DB.Find(&roles).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, roles)
}

// Handle create role
func handleCreateRole(c echo.Context) error {
	var role Role
	if err := c.Bind(&role); err != nil {
		return err
	}
	if err := DB.Create(&role).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, role)
}

// Handle update role
func handleUpdateRole(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid role id")
	}
	var role Role
	if err := DB.First(&role, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "role not found")
	}
	if err := c.Bind(&role); err != nil {
		return err
	}
	if err := DB.Save(&role).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, role)
}

// Handle delete role
func handleDeleteRole(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid role id")
	}
	var role Role
	if err := DB.First(&role, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "role not found")
	}
	if err := DB.Delete(&role).Error; err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

// Initialize routes
func RolesRoutes(e *echo.Echo) {

	roles := e.Group("/roles")
	roles.GET("", handleGetRoles)
	roles.POST("", handleCreateRole)
	roles.PUT("/:id", handleUpdateRole)
	roles.DELETE("/:id", handleDeleteRole)
}
