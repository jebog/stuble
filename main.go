package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	controller "github.com/jebog/stuble/controllers"
	"github.com/jebog/stuble/database"
	middlewares "github.com/jebog/stuble/midldlewares"
	"github.com/jebog/stuble/models"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()

	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Entry{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middlewares.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)

	err := router.Run(os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatal("Error loading using this port")
	}

	fmt.Println("Server running on port 8000")
}
