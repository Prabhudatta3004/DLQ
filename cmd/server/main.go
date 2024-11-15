package main

import (
	"fmt"
	"log"

	"github.com/Prabhudatta3004/DLQ/config"
	"github.com/Prabhudatta3004/DLQ/models"
	"github.com/Prabhudatta3004/DLQ/routers"
	"github.com/Prabhudatta3004/DLQ/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize logger
	utils.InitLogger()

	// Load configuration
	config.LoadConfig()

	// Connect to the database
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := db.AutoMigrate(&models.Message{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Setup router
	router := routers.SetupRouter(db)

	// Start server
	port := config.AppConfig.ServerPort
	fmt.Printf("Server is running on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func initDB() (*gorm.DB, error) {
	cfg := config.AppConfig
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
