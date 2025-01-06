package controllers

import (
	"buy2play/config"
	"buy2play/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// PlaceOrder places a new order for the user
func PlaceOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input struct {
		CartItems []struct {
			ProductID uint `json:"product_id"`
			Quantity  int  `json:"quantity"`
		} `json:"cart_items"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var totalPrice int
	var orderItems []models.OrderItem
	for _, item := range input.CartItems {
		if item.Quantity <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be greater than zero"})
			return
		}

		var product models.Product
		if err := config.DB.First(&product, item.ProductID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		totalPrice += product.Price * item.Quantity

		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	order := models.Order{
		UserID:     userID.(uint),
		Timestamp:  time.Now(),
		TotalPrice: totalPrice,
		Status:     models.Pending,
		OrderItems: orderItems,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order_id": order.ID, "total_price": totalPrice})
}

func GetUserOrders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	db := config.DB

	var orders []models.Order
	if err := db.Preload("OrderItems.Product").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var orderResponses []OrderResponse
	for _, order := range orders {
		var productResponses []ProductResponse
		for _, orderItem := range order.OrderItems {
			product := orderItem.Product
			productResponses = append(productResponses, ProductResponse{
				ID:          product.ID,
				Name:        product.Name,
				Price:       product.Price,
				Description: product.Description,
				ImageURL:    product.ImageURL,
				Quantity:    orderItem.Quantity,
			})
		}

		orderResponses = append(orderResponses, OrderResponse{
			OrderID:    order.ID,
			Status:     order.Status,
			Timestamp:  order.Timestamp,
			TotalPrice: order.TotalPrice,
			Products:   productResponses,
		})
	}

	c.JSON(http.StatusOK, struct {
		Orders []OrderResponse `json:"orders"`
	}{
		Orders: orderResponses,
	})
}

// GetOrder retrieves a specific order by ID
func GetOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input struct {
		OrderID int `json:"order_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB

	var order models.Order
	if err := db.Where("user_id = ?", userID).Preload("OrderItems.Product").First(&order, input.OrderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var productResponses []ProductResponse
	for _, orderItem := range order.OrderItems {
		product := orderItem.Product
		productResponses = append(productResponses, ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			ImageURL:    product.ImageURL,
			Quantity:    orderItem.Quantity,
		})
	}

	c.JSON(http.StatusOK, struct {
		OrderID    uint              `json:"order_id"`
		Status     models.Status     `json:"status"`
		Timestamp  time.Time         `json:"timestamp"`
		TotalPrice int               `json:"total_price"`
		Products   []ProductResponse `json:"products"`
	}{
		OrderID:    order.ID,
		Status:     order.Status,
		Timestamp:  order.Timestamp,
		TotalPrice: order.TotalPrice,
		Products:   productResponses,
	})
}

// UpdateOrderStatus allows admins to update the status of an order
func UpdateOrderStatus(c *gin.Context) {
	var input struct {
		OrderID uint   `json:"order_id"`
		Status  string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validStatuses := map[string]models.Status{
		string(models.Pending):   models.Pending,
		string(models.Approved):  models.Approved,
		string(models.Completed): models.Completed,
		string(models.Rejected):  models.Rejected,
	}

	status, exists := validStatuses[input.Status]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	db := config.DB
	var order models.Order
	if err := db.Preload("User").First(&order, input.OrderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.Status = status
	if err := db.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
		return
	}

	// TODO: Automatic Sender
	if status == models.Approved {
		msg := "ID Заказа: " + strconv.Itoa(int(order.ID)) + "\nТовары к выдаче:\n"
		var orderItemsManual []models.OrderItem
		for _, orderItem := range order.OrderItems {
			if orderItem.Product.Type == models.Manual {
				orderItemsManual = append(orderItemsManual, orderItem)
				msg += orderItem.Product.Name + ": " + strconv.Itoa(orderItem.Quantity) + " ед.\n"
			}
		}

		if len(orderItemsManual) > 0 {
			var admin models.User

			if err := db.Model(&models.User{}).Where("is_admin = ?", true).First(&admin).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
				return
			}

			conversation := models.Conversation{
				OrderID:   order.ID,
				Order:     order,
				UserID:    order.UserID,
				User:      order.User,
				AdminID:   admin.ID,
				Admin:     admin,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			if err := db.Save(&conversation).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Could not create conversation"})
				return
			}

			db.Save(&models.Message{
				ConversationID: conversation.ID,
				Conversation:   conversation,
				SenderID:       order.UserID,
				Sender:         order.User,
				Content:        msg,
				Timestamp:      time.Now(),
			})

			c.JSON(http.StatusOK, struct {
				Order    models.Order       `json:"order"`
				Products []models.OrderItem `json:"manual_order_items"`
			}{
				Order:    order,
				Products: orderItemsManual,
			})
			return
		}
	}

	c.JSON(http.StatusOK, struct {
		Order    models.Order       `json:"order"`
		Products []models.OrderItem `json:"manual_order_items"`
	}{
		Order:    order,
		Products: nil,
	})
}
