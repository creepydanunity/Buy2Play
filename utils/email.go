package utils

import (
	"buy2play/config"
	"buy2play/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/smtp"
	"os"
	"time"
)

func SendEmail(to string, subject string, body string) error {
	from := os.Getenv("EMAIL_SENDER")

	auth := smtp.PlainAuth(
		"",
		from,
		os.Getenv("GOOGLE_APP_PASS"),
		"smtp.gmail.com",
	)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	if err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg)); err != nil {
		return fmt.Errorf("send email to %s failed: %w", to, err)
	}
	log.Infof("Sent email from %s to %s", from, to)
	return nil
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
