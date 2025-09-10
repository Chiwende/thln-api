package routes

import (
	"github.com/gin-gonic/gin"
	"qinsights.com/thln/controllers"
)

func RegisterCallbackRoutes(router *gin.Engine) {
	router.POST("/callback", controllers.CallbackController)
}
