package controllers

import (
	"buy2play/config"
	"buy2play/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PlaceOrder places a new order for the user
func PlaceOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB
	if err := db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GetUserOrders retrieves all orders for a specific user
func GetUserOrders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	db := config.DB

	var orders []models.Order
	if err := db.Preload("Products").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// GetOrder retrieves a specific order by ID
func GetOrder(c *gin.Context) {
	orderID := c.Param("orderID")
	db := config.DB

	var order models.Order
	if err := db.Preload("Products").First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// UpdateOrderStatus allows admins to update the status of an order
func UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("orderID")
	var input struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB
	var order models.Order
	if err := db.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.Status = config.Status(input.Status)
	db.Save(&order)
	c.JSON(http.StatusOK, order)
}
