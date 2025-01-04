package controllers

import (
	"buy2play/config"
	"buy2play/models"
	"buy2play/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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

	// Hash password
	err := user.SetPassword(user.Password)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to hash password")

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Save user to DB
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

	logrus.WithFields(logrus.Fields{
		"user_id": user.ID,
	}).Info("User created successfully")
	c.JSON(http.StatusOK, gin.H{})
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

	// Retrieve user from DB (TBC)
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

	// Generate JWT token
	var token string
	if tempUser.RememberMe {
		token, err = utils.GenerateToken(user.ID, user.Username, user.Email, 14)
	} else {
		token, err = utils.GenerateToken(user.ID, user.Username, user.Email, 4)
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": "Sign In error",
		}).Error(fmt.Sprintf("Token generation error: %s", err))

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error generating token"})
		return
	}

	c.SetCookie("Authorization", token, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/")
}
