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

	err := database.Database.AutoMigrate(
		&models.User{},
		&models.UserDetails{},
		&models.Media{},
		&models.Room{},
		&models.Reservation{},
		&models.Review{},
	)
	if err != nil {
		return
	}

}

func serveApplication() {
	router := gin.Default()

	// Route
	routes.NewAuthRoute(router)
	routes.NewUserRoute(router)

	routes.NewUserDetailsRoute(router)
	routes.NewRoomRoute(router)
	routes.NewReservationRoute(router)
	routes.NewReviewRoute(router)
	routes.NewMediaRoute(router)

	err := router.Run(os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatal("Error loading using this port")
	}

	fmt.Println("Server running on port 8000")
}
