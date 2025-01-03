package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
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
