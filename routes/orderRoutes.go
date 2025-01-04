package routes

import (
	"buy2play/controllers"
	"buy2play/middlewares"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	r.POST("/orders", middlewares.AuthRequired(), controllers.PlaceOrder)
	r.GET("/orders", middlewares.AuthRequired(), controllers.GetUserOrders)
	r.GET("/orders/details/:orderID", middlewares.AuthRequired(), controllers.GetOrder)
	r.PUT("/orders/:orderID/status", middlewares.AdminAuthRequired(), controllers.UpdateOrderStatus)
}
