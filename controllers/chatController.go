package controllers

import (
	"buy2play/config"
	"buy2play/models"
	"buy2play/utils"
	"buy2play/websocketInternal"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

// GetConversations retrieves a list of all conversations for the authenticated user
func GetConversations(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var conversations []models.Conversation
	err := config.DB.Model(&models.Conversation{}).Preload("Order").Preload("User").Preload("Admin").Where("user_id = ?", userID).Find(&conversations).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve conversations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"conversations": conversations})
}

// GetConversation retrieves a particular conversation along with its messages for the authenticated user
func GetConversation(c *gin.Context) {
	conversationIDStr := c.Param("id")
	conversationID, err := strconv.Atoi(conversationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid conversation ID"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var conversation models.Conversation
	err = config.DB.Model(&models.Conversation{}).Preload("Order").Preload("User").Preload("Admin").Where("id = ? AND user_id = ?", conversationID, userID).First(&conversation).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conversation not found"})
		return
	}

	var messages []models.Message
	err = config.DB.Model(&models.Message{}).Preload("Sender").Where("conversation_id = ?", conversationID).Preload("Sender").Order("timestamp ASC").Find(&messages).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"conversation_id": conversation.ID,
		"order_id":        conversation.OrderID,
		"order":           conversation.Order,
		"user_id":         conversation.UserID,
		"user":            conversation.User,
		"admin_id":        conversation.AdminID,
		"admin":           conversation.Admin,
		"created_at":      conversation.CreatedAt,
		"updated_at":      conversation.UpdatedAt,
		"messages":        messages,
	})
}

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins TODO: Change for prod
	},
}

// WebSocketHandler handles WebSocket communication for a specific conversation
func WebSocketHandler(c *gin.Context) {
	token, err := c.Cookie("Authorization")
	if err != nil || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	claims, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.Set("userID", claims.ID)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Error("Failed to upgrade connection:", err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Error("Failed to close websocketInternal connection", err)
		}
	}(conn)

	var input struct {
		OrderID int `json:"order_id"`
	}

	input.OrderID, _ = strconv.Atoi(c.Param("orderID"))

	// Get userID from context (ensure user is authenticated)
	userID, exists := c.Get("userID")
	if !exists {
		err := conn.WriteMessage(websocket.TextMessage, []byte("User not authenticated"))
		if err != nil {
			log.Error("Failed to write message:", err)
			return
		}
		return
	}

	var conversation models.Conversation
	err = config.DB.Where("order_id = ? AND user_id = ?", input.OrderID, userID).Preload("Order").First(&conversation).Error
	if err != nil {
		err := conn.WriteMessage(websocket.TextMessage, []byte("Unauthorized access to conversation"))
		if err != nil {
			log.Error("Failed to write message:", err)
			return
		}
		return
	}

	websocketInternal.AddClient(uint(input.OrderID), conn)

	err = conn.WriteMessage(websocket.TextMessage, []byte("Вы зашли в чат!"))
	if err != nil {
		log.Error("Failed to write message:", err)
		return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Error("Error reading message:", err)
			break
		}

		message := models.Message{
			ConversationID: conversation.ID,
			SenderID:       userID.(uint),
			Content:        string(msg),
			Timestamp:      time.Now(),
		}

		err = config.DB.Create(&message).Error
		if err != nil {
			err = conn.WriteMessage(websocket.TextMessage, []byte("Failed to save message"))
			if err != nil {
				log.Error("Failed to write message:", err)
			}
			continue
		}

		websocketInternal.BroadcastMessage(uint(input.OrderID), msg)
	}
	websocketInternal.RemoveClient(uint(input.OrderID), conn)
}
