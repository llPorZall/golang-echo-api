package api

import (
	"api/customer"
	"api/db"
	_ "api/docs/api"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//StartServer function for start echo server
func StartServer() {
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Error loading .env file")
	}
	fmt.Println("------------ GOLANG ECHO API SERVER START -------------")

	router := echo.New()

	// ===== Middlewares
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	// ===== Prefix of all routes
	publicRoute := router.Group("/api/v1")
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	// ===== Initial resource from MongoDB
	resource, err := db.CreateResource()
	if err != nil {
		logrus.Error(err)
	}
	defer resource.Close()

	// ===== Add routes of users
	customer.CustomerAPI(publicRoute, resource)

	// ===== Start server
	router.Start(":8080")
}
