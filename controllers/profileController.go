package controllers

import (
	"buy2play/config"
	"buy2play/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type OrderResponse struct {
	OrderID    uint              `json:"order_id"`
	Timestamp  time.Time         `json:"timestamp"`
	TotalPrice int               `json:"total_price"`
	Products   []ProductResponse `json:"products"`
}

type ProductResponse struct {
	ID          uint   `json:"product_id"`
	Name        string `json:"product_name"`
	Price       int    `json:"product_price"`
	Description string `json:"product_description"`
	ImageURL    string `json:"product_image_url"`
}

// GetUserProfile retrieves a user's profile
func GetUserProfile(c *gin.Context) {
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

	var userOrders []models.Order
	if err := db.Model(&models.Order{}).Where("user_id = ?", user.ID).
		Preload("Products").
		Select("id", "timestamp", "total_price", "status").
		Find(&userOrders).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User orders not found"})
	}

	var orderResponses []OrderResponse
	for _, order := range userOrders {
		var productResponses []ProductResponse
		for _, product := range order.Products {
			productResponses = append(productResponses, ProductResponse{
				ID:          product.ID,
				Name:        product.Name,
				Price:       product.Price,
				Description: product.Description,
				ImageURL:    product.ImageURL,
			})
		}

		orderResponses = append(orderResponses, OrderResponse{
			OrderID:    order.ID,
			Timestamp:  order.Timestamp,
			TotalPrice: order.TotalPrice,
			Products:   productResponses,
		})
	}

	c.JSON(http.StatusOK, struct {
		Email    string          `json:"email"`
		Username string          `json:"username"`
		Orders   []OrderResponse `json:"orders"`
	}{
		Email:    user.Email,
		Username: user.Username,
		Orders:   orderResponses,
	})
}

type profileEditInput struct {
	Username string `json:"username"`
}

// UpdateUserProfile updates a user's profile
func UpdateUserProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input profileEditInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Username = input.Username
	db.Save(user)

	c.JSON(http.StatusOK, struct {
		NewName string `json:"username"`
	}{
		NewName: user.Username,
	})
}
