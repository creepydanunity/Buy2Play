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
	BaseURL   string
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
	BaseURL = os.Getenv("BASE_URL")
}

func SyncDatabase() {
	if DB.AutoMigrate(&models.User{}) != nil {
		panic("Database models.User migration failed")
	}

	if DB.AutoMigrate(&models.ProductCategory{}) != nil {
		panic("Database models.ProductCategory migration failed")
	}

	if DB.AutoMigrate(&models.ProductSubCategory{}) != nil {
		panic("Database models.ProductSubCategory migration failed")
	}

	if DB.AutoMigrate(&models.Product{}) != nil {
		panic("Database models.Product migration failed")
	}

	if DB.AutoMigrate(&models.VerificationToken{}) != nil {
		panic("Database models.VerificationToken migration failed")
	}

	if DB.AutoMigrate(&models.CartItem{}) != nil {
		panic("Database models.CartItem migration failed")
	}

	if DB.AutoMigrate(&models.OrderItem{}) != nil {
		panic("Database models.OrderItem migration failed")
	}

	if DB.AutoMigrate(&models.Order{}) != nil {
		panic("Database models.Order migration failed")
	}

	if DB.AutoMigrate(&models.Conversation{}) != nil {
		panic("Database models.Conversation migration failed")
	}

	if DB.AutoMigrate(&models.Message{}) != nil {
		panic("Database models.Message migration failed")
	}
}
