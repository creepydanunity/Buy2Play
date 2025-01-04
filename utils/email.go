package utils

import (
	"buy2play/config"
	"buy2play/models"
	"errors"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"os"
	"time"
)

func SendEmail(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_SENDER"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		587,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	return d.DialAndSend(m)
}

func GenerateVerificationToken() string {
	return uuid.New().String() + "-" + uuid.New().String()
}

func SaveVerificationToken(db *gorm.DB, userID uint, verificationToken string) error {
	expiryTime := time.Now().Add(10 * time.Minute)

	token := models.VerificationToken{
		Token:     verificationToken,
		UserID:    userID,
		CreatedAt: time.Now(),
		ExpiresAt: expiryTime,
	}

	if err := db.Create(&token).Error; err != nil {
		return errors.New("failed to save verification token: " + err.Error())
	}

	return nil
}

func VerifyToken(token string) (uint, error) {
	var verificationToken models.VerificationToken
	db := config.DB

	if err := db.Where("token = ? AND expires_at > ? AND used = ?", token, time.Now(), false).First(&verificationToken).Error; err != nil {
		return 0, errors.New("invalid or expired token")
	}

	verificationToken.Used = true
	db.Save(&verificationToken)

	return verificationToken.UserID, nil
}
