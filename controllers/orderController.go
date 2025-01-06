package controllers

import (
	"buy2play/config"
	"buy2play/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type ProductResponseTemp struct {
	ProductID    uint               `json:"product_id"`
	ProductName  string             `json:"product_name"`
	ProductPrice int                `json:"product_price"`
	ProductType  models.ProductType `json:"product_type"`
	ProductImage string             `json:"product_image_url"`
}

type OrderItemResponse struct {
	OrderItemID uint                `json:"order_item_id"`
	OrderID     uint                `json:"order_id"`
	Quantity    int                 `json:"quantity"`
	Product     ProductResponseTemp `json:"product"`
}

type OrderResponseTemp struct {
	OrderID    uint                `json:"order_id"`
	Timestamp  time.Time           `json:"timestamp"`
	TotalPrice int                 `json:"total_price"`
	OrderItems []OrderItemResponse `json:"order_items"`
	Status     models.Status       `json:"status"`
}

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
	if err := db.Preload("OrderItems.Product").First(&order, input.OrderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	order.Status = status
	if err := db.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
		return
	}

	orderResponse := transformOrder(order)

	var manualOrderItems []OrderItemResponse
	for _, item := range order.OrderItems {
		if item.Product.Type == models.Manual {
			manualOrderItems = append(manualOrderItems, transformOrderItem(item))
		}
	}
	if status == models.Approved {
		msg := "ID Заказа: " + strconv.Itoa(int(order.ID)) + "\nТовары к выдаче:\n"
		for _, orderItem := range order.OrderItems {
			if orderItem.Product.Type == models.Manual {
				manualOrderItems = append(manualOrderItems, transformOrderItem(orderItem))
				msg += orderItem.Product.Name + ": " + strconv.Itoa(orderItem.Quantity) + " ед.\n"
			}
		}

		if len(manualOrderItems) > 0 {
			var admin models.User

			if err := db.Model(&models.User{}).Where("is_admin = ?", true).First(&admin).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Admin not found"})
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
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create conversation", "details": err})
				return
			}

			if err := db.Save(&models.Message{
				ConversationID: conversation.ID,
				Conversation:   conversation,
				SenderID:       order.UserID,
				Sender:         order.User,
				Content:        msg,
				CreatedAt:      time.Now(),
			}); err != nil {
				log.Errorf("Could not create initial message: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create message", "details": err})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"order":              orderResponse,
				"manual_order_items": manualOrderItems,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"order":              orderResponse,
		"manual_order_items": manualOrderItems,
	})
}

func transformOrder(order models.Order) OrderResponseTemp {
	var orderItems []OrderItemResponse
	for _, item := range order.OrderItems {
		orderItems = append(orderItems, transformOrderItem(item))
	}
	return OrderResponseTemp{
		OrderID:    order.ID,
		Timestamp:  order.Timestamp,
		TotalPrice: order.TotalPrice,
		OrderItems: orderItems,
		Status:     order.Status,
	}
}

func transformOrderItem(item models.OrderItem) OrderItemResponse {
	return OrderItemResponse{
		OrderItemID: item.ID,
		OrderID:     item.OrderID,
		Quantity:    item.Quantity,
		Product: ProductResponseTemp{
			ProductID:    item.Product.ID,
			ProductName:  item.Product.Name,
			ProductPrice: item.Product.Price,
			ProductType:  item.Product.Type,
			ProductImage: item.Product.ImageURL,
		},
	}
}
