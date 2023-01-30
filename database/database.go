package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Database *gorm.DB

func Connect() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")


	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
    
    retries := 5
	for i := 0; i < retries; i++ { 
        
		// Open connection to database
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
		Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger:      newLogger,
			PrepareStmt: true,
			QueryFields: true, // Select by all fields' name
		})

		if err != nil {
			if i == retries-1 {
				log.Fatal("Error connecting to database with error: ", err)
			}
			log.Printf("Error connecting to database, retrying... (%v/%v)\n",i+1,retries)
			// Wait for 1 second before retrying
			time.Sleep(time.Second)
		} else {
			log.Println("Connected to database at " + host + ":" + port)
			break
		}
    }
}
