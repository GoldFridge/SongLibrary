package database

import (
	"SongLibrary/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Открытие соединения с базой данных через GORM
	DB, err = gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONN")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(models.Song{})
	if err != nil {
		log.Println("Failed to migrate database:", err)
	}
}
