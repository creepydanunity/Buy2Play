package config

import (
	"buy2play/models"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	JWTSecret []byte
	DB        *gorm.DB
)

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Rejected Status = "rejected"
)

type ProductType string

const (
	Auto   ProductType = "auto"
	Manual ProductType = "manual"
)

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Fail to connect to DB")
	}
}

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
}

func SyncDatabase() {
	if DB.AutoMigrate(&models.User{}) != nil {
		panic("Database models.User migration failed")
	}
}
