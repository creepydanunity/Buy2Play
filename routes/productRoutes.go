package routes

import (
	"buy2play/controllers"
	"buy2play/middlewares"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/products", middlewares.AdminAuthRequired(), controllers.GetAllProducts)
	r.GET("/products/categories", controllers.GetCategories)
	r.GET("/products/categorized", controllers.GetCategoryProducts)
	r.GET("/products/product", controllers.GetProduct)
	r.POST("/products/product", middlewares.AdminAuthRequired(), controllers.CreateProduct)
	r.PATCH("/products/product", middlewares.AdminAuthRequired(), controllers.UpdateProduct)
	r.DELETE("/products/product", middlewares.AdminAuthRequired(), controllers.DeleteProduct)
}
