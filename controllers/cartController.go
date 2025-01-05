package controllers

import (
	"buy2play/config"
	"buy2play/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type cartItemInput struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

// AddItemToCart adds an item to the user's cart
func AddItemToCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var cartInput cartItemInput
	if err := c.ShouldBindJSON(&cartInput); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": "Failed to read body",
		}).Error("AddItemToCart error")

		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	if cartInput.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be greater than zero"})
		return
	}

	db := config.DB

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var product models.Product
	if err := db.First(&product, cartInput.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var existingCartItem models.CartItem
	err := db.Where("product_id = ? AND user_id = ?", cartInput.ProductID, userID).First(&existingCartItem).Error
	if err == nil {
		existingCartItem.Quantity += cartInput.Quantity
		db.Save(&existingCartItem)
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Added %d units of the product", cartInput.Quantity)})
		return
	}

	var cartItem models.CartItem
	cartItem.ProductID = product.ID
	cartItem.Product = product
	cartItem.UserID = user.ID
	cartItem.User = user
	cartItem.Quantity = cartInput.Quantity

	if err := db.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Added %d units of the product", cartInput.Quantity)})
}

// GetCartResponse represents the structure of the response for a cart item
type GetCartResponse struct {
	CartID   uint                   `json:"cart_id"`
	Quantity int                    `json:"quantity"`
	Product  GetCartProductResponse `json:"product"`
}

// GetCartProductResponse represents the product structure in the cart response
type GetCartProductResponse struct {
	ProductID          uint   `json:"product_id"`
	ProductName        string `json:"product_name"`
	ProductPrice       int    `json:"product_price"`
	ProductDescription string `json:"product_description"`
	ProductType        string `json:"product_type"`
	ProductImageURL    string `json:"product_image_url"`
}

// GetCart fetches the user's cart and formats the response
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

	if len(cartItems) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Your cart is empty"})
		return
	}

	// Transform cart items into the desired response format
	var response []GetCartResponse
	for _, item := range cartItems {
		response = append(response, GetCartResponse{
			CartID:   item.ID,
			Quantity: item.Quantity,
			Product: GetCartProductResponse{
				ProductID:          item.Product.ID,
				ProductName:        item.Product.Name,
				ProductPrice:       item.Product.Price,
				ProductDescription: item.Product.Description,
				ProductType:        string(item.Product.Type),
				ProductImageURL:    item.Product.ImageURL,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}

// UpdateCartItem updates the quantity of an item in the cart based on product ID
func UpdateCartItem(c *gin.Context) {
	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if input.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be greater than zero"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	db := config.DB

	var cartItem models.CartItem
	if err := db.Where("product_id = ? AND user_id = ?", input.ProductID, userID).First(&cartItem).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		return
	}

	cartItem.Quantity = input.Quantity
	if err := db.Save(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
		return
	}

	response := GetCartResponse{
		CartID:   cartItem.ID,
		Quantity: cartItem.Quantity,
		Product: GetCartProductResponse{
			ProductID:          cartItem.Product.ID,
			ProductName:        cartItem.Product.Name,
			ProductPrice:       cartItem.Product.Price,
			ProductDescription: cartItem.Product.Description,
			ProductType:        string(cartItem.Product.Type),
			ProductImageURL:    cartItem.Product.ImageURL,
		},
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart item updated", "item": response})
}

// RemoveItemFromCart removes an item from the user's cart based on product ID
func RemoveItemFromCart(c *gin.Context) {
	var input struct {
		ProductID uint `json:"product_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	db := config.DB

	var cartItem models.CartItem
	if err := db.Where("product_id = ? AND user_id = ?", input.ProductID, userID).First(&cartItem).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		return
	}

	if err := db.Delete(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
