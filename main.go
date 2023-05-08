package main

import (
	"log"
	"os"
	"projet-go/handlers"
	"projet-go/models"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Main function
func main() {

	// Initialize environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Open a connection to the database
	models.DB, err = models.SetupDB()
	if err != nil {
		panic("err")
	}

	// Create an instance of the Echo framework
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Init Routes
	models.UsersRoutes(e)
	models.GroupsRoutes(e)
	models.RolesRoutes(e)
	models.AuthsRoutes(e)

	// Initialize JWT secret
	//jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Initialize middleware
	// jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: jwtSecret,
	// })

	// Initialize handlers
	handlers.UsersHandler(e)
	handlers.GroupsHandler(e)
	handlers.RolesHandler(e)
	handlers.AuthsHandler(e)

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
