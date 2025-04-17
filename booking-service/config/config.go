package config

import (
	"booking-service/model"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitConfig() {
	InitLogging()
	InitEnv()
	InitDB()

}

var DB *gorm.DB

func InitDB() {
	log.Info("🔧 Config DB")
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&model.Booking{})
	if err != nil {
		return
	}
	DB = db
}

func InitLogging() {
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	log.SetReportCaller(true)
}

func InitEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}
}
