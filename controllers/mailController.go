package controllers

import (
	"buy2play/config"
	"buy2play/models"
	"buy2play/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SendVerificationEmail sends a verification email to the user
func SendVerificationEmail(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	db := config.DB

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.EmailVerified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email is already verified"})
		return
	}

	if !user.LastVerificationSent.IsZero() && time.Since(user.LastVerificationSent) < 5*time.Minute {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Please wait before requesting another verification email"})
		return
	}

	verificationToken := utils.GenerateVerificationToken()
	link := config.BaseURL + "/verify-email?token=" + verificationToken

	if err := utils.SaveVerificationToken(db, user.ID, verificationToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save verification token", "details": err.Error()})
		return
	}

	emailBody := "Click <a href='" + link + "'>here</a> to verify your email."
	if err := utils.SendEmail(user.Email, "Email Verification", emailBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db.Model(&user).Update("LastVerificationSent", time.Now())

	c.JSON(http.StatusOK, gin.H{"message": "Verification email sent"})
}

// VerifyEmail verifies the user's email
func VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	db := config.DB

	userID, err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired token"})
		return
	}

	db.Model(&models.User{}).Where("id = ?", userID).Update("email_verified", true)
	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}
