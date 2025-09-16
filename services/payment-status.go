package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// PaymentStatusUpdateRequest represents the request structure for updating payment status
type PaymentStatusUpdateRequest struct {
	OrderID       string `json:"orderId"`
	PaymentStatus string `json:"paymentStatus"`
	TransactionID string `json:"transactionId"`
	Notes         string `json:"notes"`
}

// PaymentStatusUpdateResponse represents the response from the payment status update API
type PaymentStatusUpdateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		OrderID       string `json:"orderId"`
		PaymentStatus string `json:"paymentStatus"`
		UpdatedAt     string `json:"updatedAt"`
	} `json:"data,omitempty"`
}

// UpdatePaymentStatus sends a PUT request to update the payment status
func UpdatePaymentStatus(orderID, paymentStatus, transactionID, notes string) (*PaymentStatusUpdateResponse, error) {
	// Get the API base URL from environment variables
	apiBaseURL := os.Getenv("API_BASE_URL")
	if apiBaseURL == "" {
		apiBaseURL = "https://your-domain.com" // fallback URL
	}

	// Construct the full URL
	url := fmt.Sprintf("%s/api/orders/payment-status", apiBaseURL)

	// Create the request payload
	request := PaymentStatusUpdateRequest{
		OrderID:       orderID,
		PaymentStatus: paymentStatus,
		TransactionID: transactionID,
		Notes:         notes,
	}

	// Marshal the request to JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		log.Printf("Error marshaling payment status update request: %v", err)
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Set headers
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	// Send the PUT request
	responseBody, statusCode, err := SendRequest("PUT", url, headers, requestBody)
	if err != nil {
		log.Printf("Error sending payment status update request: %v", err)
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// Check if the request was successful
	if statusCode < 200 || statusCode >= 300 {
		log.Printf("Payment status update failed with status code: %d", statusCode)
		return nil, fmt.Errorf("payment status update failed with status code: %d", statusCode)
	}

	// Parse the response
	var response PaymentStatusUpdateResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		log.Printf("Error unmarshaling payment status update response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	log.Printf("Payment status updated successfully for order %s", orderID)
	return &response, nil
}

// MapCallbackStatusToPaymentStatus maps the callback status to payment status
func MapCallbackStatusToPaymentStatus(callbackStatus string) string {
	switch callbackStatus {
	case "success":
		return "COMPLETED"
	case "failed":
		return "FAILED"
	case "pending":
		return "PENDING"
	default:
		return "UNKNOWN"
	}
}
