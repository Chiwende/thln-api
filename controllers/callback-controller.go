package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	services "qinsights.com/thln/services"
)

// CallbackResponse represents the structure of the callback response
type CallbackResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		TransactionReference string `json:"transaction_reference"`
		ExternalReference    string `json:"external_reference"`
		Customer             string `json:"customer"`
		Amount               string `json:"amount"`
	} `json:"data"`
}

func CallbackController(c *gin.Context) {
	var callbackResponse CallbackResponse

	// Bind the JSON request body to the struct
	if err := c.ShouldBindJSON(&callbackResponse); err != nil {
		log.Printf("Error binding callback JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Log the structured callback response
	log.Printf("Callback received - Code: %d, Status: %s, Message: %s",
		callbackResponse.Code, callbackResponse.Status, callbackResponse.Message)
	log.Printf("Transaction Details - Reference: %s, External: %s, Customer: %s, Amount: %s",
		callbackResponse.Data.TransactionReference,
		callbackResponse.Data.ExternalReference,
		callbackResponse.Data.Customer,
		callbackResponse.Data.Amount)

	// Update payment status via external API
	paymentStatus := services.MapCallbackStatusToPaymentStatus(callbackResponse.Status)
	notes := fmt.Sprintf("Payment %s by gateway - %s", callbackResponse.Status, callbackResponse.Message)

	// Use external reference as order ID, or transaction reference as fallback
	orderID := callbackResponse.Data.ExternalReference
	if orderID == "" {
		orderID = callbackResponse.Data.TransactionReference
	}

	updateResponse, err := services.UpdatePaymentStatus(
		orderID,
		paymentStatus,
		callbackResponse.Data.TransactionReference,
		notes,
	)

	if err != nil {
		log.Printf("Error updating payment status: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Callback received but failed to update payment status",
			"error":   err.Error(),
		})
		return
	}

	log.Printf("Payment status updated successfully: %+v", updateResponse)
	c.JSON(http.StatusOK, gin.H{
		"message":        "Callback received and payment status updated",
		"updateResponse": updateResponse,
	})
}
