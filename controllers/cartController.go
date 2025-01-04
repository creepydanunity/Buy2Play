package controllers

import (
	"buy2play/config"
	"buy2play/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddItemToCart adds an item to the user's cart
func AddItemToCart(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB

	var product models.Product
	if err := db.First(&product, cartItem.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var existingCartItem models.CartItem
	err := db.Where("product_id = ? AND user_id = ?", cartItem.ProductID, cartItem.UserID).First(&existingCartItem).Error
	if err == nil {
		existingCartItem.Quantity += cartItem.Quantity
		db.Save(&existingCartItem)
		c.JSON(http.StatusOK, existingCartItem)
		return
	}

	cartItem.Product = product
	db.Create(&cartItem)
	c.JSON(http.StatusCreated, cartItem)
}

// GetCart fetches the user's cart
func GetCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	db := config.DB

	var cartItems []models.CartItem
	if err := db.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

// UpdateCartItem updates the quantity of a cart item
func UpdateCartItem(c *gin.Context) {
	cartID := c.Param("cartID")
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB
	var existingCartItem models.CartItem
	if err := db.First(&existingCartItem, cartID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		return
	}

	existingCartItem.Quantity = cartItem.Quantity
	db.Save(&existingCartItem)
	c.JSON(http.StatusOK, existingCartItem)
}

// RemoveItemFromCart removes an item from the user's cart
func RemoveItemFromCart(c *gin.Context) {
	cartID := c.Param("cartID")
	db := config.DB

	if err := db.Delete(&models.CartItem{}, cartID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
