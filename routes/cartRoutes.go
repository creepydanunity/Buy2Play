package routes

import (
	"buy2play/controllers"
	"buy2play/middlewares"
	"github.com/gin-gonic/gin"
)

func CartRoutes(r *gin.Engine) {
	r.POST("/cart", middlewares.AuthRequired(), controllers.AddItemToCart)
	r.GET("/cart", middlewares.AuthRequired(), controllers.GetCart)
	r.PATCH("/cart", middlewares.AuthRequired(), controllers.UpdateCartItem)
	r.DELETE("/cart", middlewares.AuthRequired(), controllers.RemoveItemFromCart)
}
