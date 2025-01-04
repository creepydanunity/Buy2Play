package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID                   uint      `json:"user_id"`
	Email                string    `json:"email"`
	Username             string    `json:"username"`
	Password             string    `json:"password"`
	EmailVerified        bool      `json:"email_verified" gorm:"default:false"`
	IsAdmin              bool      `json:"is_admin" gorm:"default:false"`
	LastVerificationSent time.Time `json:"last_verification_sent"`
}

// SetPassword Hash password
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword Validate password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
