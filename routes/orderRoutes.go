package routes

import (
	"buy2play/controllers"
	"buy2play/middlewares"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	r.POST("/orders", middlewares.AuthRequired(), controllers.PlaceOrder)
	r.GET("/orders", middlewares.AuthRequired(), controllers.GetUserOrders)
	r.GET("/orders/details", middlewares.AuthRequired(), controllers.GetOrder)
	r.PATCH("/orders/status", middlewares.AdminAuthRequired(), controllers.UpdateOrderStatus)
}
