package models

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// Group model
type Group struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"uniqueIndex;not null"`
	Parent_Group_ID uint      `json:"parent_group_id" gorm:"default:null"`
	Child_Group_IDs []*uint   `json:"child_group_ids" gorm:"-"`
	Created_At      time.Time `json:"created_at" gorm:"not null"`
	Updated_At      time.Time `json:"updated_at" gorm:"not null"`
	Deleted_At      time.Time `json:"deleted_at" gorm:"index"`
	Users           []*User   `json:"users" gorm:"many2many:user_groups;"`
	Roles           []*Role   `json:"roles" gorm:"many2many:group_roles;"`
}

// Handle get groups
func handleGetGroups(c echo.Context) error {
	var groups []Group
	if err := DB.Find(&groups).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, groups)
}

// Handle create group
func handleCreateGroup(c echo.Context) error {
	var group Group
	if err := c.Bind(&group); err != nil {
		return err
	}
	if err := DB.Create(&group).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, group)
}

// Handle update group
func handleUpdateGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid group id")
	}
	var group Group
	if err := DB.First(&group, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "group not found")
	}
	if err := c.Bind(&group); err != nil {
		return err
	}
	if err := DB.Save(&group).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, group)
}

// Handle delete group
func handleDeleteGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid group id")
	}
	var group Group
	if err := DB.First(&group, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "group not found")
	}
	if err := DB.Delete(&group).Error; err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)

}

// Initialize routes
func GroupsRoutes(e *echo.Echo) {

	groups := e.Group("/groups")
	groups.GET("", handleGetGroups)
	groups.POST("", handleCreateGroup)
	groups.PUT("/:id", handleUpdateGroup)
	groups.DELETE("/:id", handleDeleteGroup)
}
