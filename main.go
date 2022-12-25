package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/database"
	"github.com/jebog/stuble/models"
	"github.com/jebog/stuble/routes"
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
	database.Database.AutoMigrate(&models.UserDetails{})
	database.Database.AutoMigrate(&models.Media{})
	database.Database.AutoMigrate(&models.Room{})
	database.Database.AutoMigrate(&models.Reservation{})
	database.Database.AutoMigrate(&models.Review{})

}

func serveApplication() {
	router := gin.Default()

	// Route
	routes.NewAuthRoute(router)

	routes.NewUserRoute(router)
	routes.NewUserDetailsRoute(router)
	routes.NewRoomRoute(router)

	err := router.Run(os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatal("Error loading using this port")
	}

	fmt.Println("Server running on port 8000")
}
