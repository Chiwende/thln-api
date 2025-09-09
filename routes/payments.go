package routes

import (
	"github.com/gin-gonic/gin"
	"qinsights.com/thln/controllers"
)

func RegisterPaymentRoutes(router *gin.Engine) {
	router.POST("/collection", controllers.GeePayCollection)
	router.GET("/transaction-status", controllers.GeePayTransactionStatus)
}
