package routes

import (
	"buy2play/controllers"
	"buy2play/middlewares"
	"github.com/gin-gonic/gin"
)

func MailRoutes(r *gin.Engine) {
	auth := r.Group("/auth", middlewares.AuthRequired())

	auth.POST("/send-verification-email", controllers.SendVerificationEmail)

	r.GET("/verify-email", controllers.VerifyEmail)
}
