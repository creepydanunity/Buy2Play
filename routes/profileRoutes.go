package routes

import (
	"buy2play/controllers"
	"buy2play/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/users", middlewares.AuthRequired(), controllers.GetUserProfile)
	r.PATCH("/users", middlewares.AuthRequired(), controllers.UpdateUserProfile)
}
