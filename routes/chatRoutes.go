package routes

import (
	"buy2play/controllers"
	"buy2play/middlewares"
	"github.com/gin-gonic/gin"
)

// ChatRoutes sets up the routes for conversation and message management
func ChatRoutes(r *gin.Engine) {
	r.GET("/conversations", middlewares.AuthRequired(), controllers.GetConversations)

	r.GET("/conversations/:id", middlewares.AuthRequired(), controllers.GetConversation)

	r.GET("/ws/conversations/:orderID", controllers.WebSocketHandler)
}
