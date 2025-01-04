package routes

import (
	"buy2play/controllers"
	"buy2play/middlewares"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetAllProducts)
	r.GET("/products/:id", controllers.GetProduct)
	r.POST("/products", middlewares.AdminAuthRequired(), controllers.CreateProduct)
	r.PUT("/products/:id", middlewares.AdminAuthRequired(), controllers.UpdateProduct)
	r.DELETE("/products/:id", middlewares.AdminAuthRequired(), controllers.DeleteProduct)
}
