package routes

import (
	"github.com/gin-gonic/gin"
	"qinsights.com/thln/controllers"
)

func RegisterKYCRoutes(router *gin.Engine) {
	router.GET("/kyc/:phone", controllers.KYCController)
}
