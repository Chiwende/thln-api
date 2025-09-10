package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	services "qinsights.com/thln/services/gee-pay"
)

func GeePayCollection(c *gin.Context) {
	var request services.GeePayCollectionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := services.GeePayCollection(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func GeePayTransactionStatus(c *gin.Context) {
	transactionRef := c.Query("transaction_ref")
	log.Println("Transaction reference:", transactionRef)
	if transactionRef == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "transaction_ref query parameter is required"})
		return
	}

	response, err := services.GeePayTransactionEnquiry(transactionRef)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
