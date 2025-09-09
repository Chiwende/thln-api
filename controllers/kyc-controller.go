package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	services "qinsights.com/thln/services/gee-pay"
)

func KYCController(c *gin.Context) {
	phone := c.Params.ByName("phone")
	log.Println("Phone:", phone)
	response, errorResponse, err := services.KYCService(phone)
	log.Println("Response:", response)
	log.Println("Error response:", errorResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
