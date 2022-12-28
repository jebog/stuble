package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jebog/stuble/database"
	_ "github.com/jebog/stuble/docs"
	"github.com/jebog/stuble/models"
	"github.com/jebog/stuble/routes"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// @title           Stuble API
// @version         1.0
// @description     API for House Rent
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  me@marclabs.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth
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
	r := gin.Default()

	// Route

	routes.NewAuthRoute(r)
	routes.NewUserRoute(r)

	routes.NewUserDetailsRoute(r)
	routes.NewRoomRoute(r)
	routes.NewReservationRoute(r)
	routes.NewReviewRoute(r)
	routes.NewMediaRoute(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatal("Error loading using this port")
	}

	fmt.Println("Server running on port 8000")
}
