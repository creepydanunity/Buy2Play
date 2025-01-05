package controllers

import (
	"buy2play/config"
	"buy2play/models"
	"buy2play/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type signUpBody struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInBody struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe"`
}

func Register(c *gin.Context) {
	var tempUser signUpBody
	if err := c.ShouldBindJSON(&tempUser); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": "Failed to read body",
		}).Error("Signup error")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	user := models.User{Email: tempUser.Email, Username: tempUser.Name}

	err := user.SetPassword(tempUser.Password)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to hash password")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	result := config.DB.Create(&user)

	if result.Error != nil {
		logrus.WithFields(logrus.Fields{
			"error": result.Error.Error(),
		}).Error("Failed to create user")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	go func() {
		verificationToken := utils.GenerateVerificationToken()
		link := config.BaseURL + "/verify-email?token=" + verificationToken

		if err := utils.SaveVerificationToken(config.DB, user.ID, verificationToken); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save verification token", "details": err.Error()})
			return
		}

		emailBody := "Перейдите по ссылке: " + link + " для завершения создания аккаунта."
		if err := utils.SendEmail(user.Email, "Buy2Play: Подтверждение почты", emailBody); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		config.DB.Model(&user).Update("LastVerificationSent", time.Now())
	}()
	defer func() {
		logrus.WithFields(logrus.Fields{
			"user_id": user.ID,
		}).Info("User created successfully. Verification sent")
		c.JSON(http.StatusOK, gin.H{"message": "Verification email sent"})
	}()
}

func Login(c *gin.Context) {
	var tempUser signInBody
	if err := c.ShouldBindJSON(&tempUser); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": "Failed to read body",
		}).Error("Signup error")

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to read body.",
		})

		return
	}

	var user models.User
	err := config.DB.Model(models.User{}).Where("email = ?", tempUser.Email).First(&user).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": "Sign In error",
		}).Error(err)

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Неверное имя пользователя или пароль",
		})

		return
	}

	if !user.CheckPassword(tempUser.Password) {
		logrus.WithFields(logrus.Fields{
			"error": "Sign In error",
		}).Error("Неверное имя пользователя или пароль")

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Неверное имя пользователя или пароль",
		})

		return
	}

	if !user.EmailVerified {
		logrus.WithFields(logrus.Fields{
			"error":  "Sign In error",
			"status": "Email not verified",
		})

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Почта не подтверждена",
		})
		return
	}

	var token string
	if tempUser.RememberMe {
		token, err = utils.GenerateToken(user.ID, user.Username, user.Email, user.IsAdmin, 14)
	} else {
		token, err = utils.GenerateToken(user.ID, user.Username, user.Email, user.IsAdmin, 4)
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": "Sign In error",
		}).Error(fmt.Sprintf("Token generation error: %s", err))

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error generating token"})
		return
	}

	c.SetCookie("Authorization", token, 3600, "/", "localhost", false, true)
	logrus.WithFields(logrus.Fields{
		"user_id": user.ID,
	}).Info("User logged successfully")
	c.JSON(http.StatusOK, gin.H{})
}

func Logout(c *gin.Context) {
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)
	logrus.WithFields(logrus.Fields{
		"user_id": userId,
	}).Info("User logged out successfully")
	c.JSON(http.StatusOK, gin.H{})
}
