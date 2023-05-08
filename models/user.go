package models

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	ID          uint         `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"not null"`
	Email       string       `json:"email" gorm:"uniqueIndex;not null"`
	Password    string       `json:"password" gorm:"not null"`
	Roles       []*Role      `json:"roles" gorm:"many2many:user_roles;"`
	Groups      []*Group     `json:"groups" gorm:"many2many:user_groups;"`
	Created_At  time.Time    `json:"created_at" gorm:"not null"`
	Updated_At  time.Time    `json:"updated_at" gorm:"not null"`
	Deleted_At  time.Time    `json:"deleted_at" gorm:"index"`
	Auth_Tokens []*AuthToken `json:"auth_tokens" gorm:"foreignKey:UserID"`
}

// Handle get users
func handleGetUsers(c echo.Context) error {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

// Handle create user
func handleCreateUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Handle update user
func handleUpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := DB.Save(&user).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Handle delete user
func handleDeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	if err := DB.Delete(&user).Error; err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

// Initialize routes
func UsersRoutes(e *echo.Echo) {
	users := e.Group("/users")
	users.GET("", handleGetUsers)
	users.POST("", handleCreateUser)
	users.PUT("/:id", handleUpdateUser)
	users.DELETE("/:id", handleDeleteUser)
}
