package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	services "qinsights.com/thln/services"
)

type GeePayAuthRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type GeePayAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type GeePayCollectionRequest struct {
	Amount      int    `json:"amount"`
	PhoneNumber string `json:"phone_number"`
}

type CollectionResponse struct {
	Code string `json:"code"`
	Data struct {
		ExternalReference    string `json:"external_reference"`
		TransactionReference string `json:"transaction_reference"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type TransactionStatusResponse struct {
	Code string `json:"code"`
	Data struct {
		Amount               string `json:"amount"`
		Status               string `json:"status"`
		TransactionReference string `json:"transaction_reference"`
		ExternalReference    string `json:"external_reference"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

const (
	HeaderContentType    = "Content-Type"
	HeaderAccept         = "Accept"
	HeaderClientID       = "X-Client-ID"
	HeaderAuthorization  = "Authorization"
	HeaderTransactionRef = "X-Transaction-Ref"
	HeaderCallbackURL    = "X-CALLBACK-URL"
)

func GeePayGenerateToken() (*GeePayAuthResponse, error) {
	data := GeePayAuthRequest{
		ClientID:     os.Getenv("GEE_PAY_CLIENT_ID"),
		ClientSecret: os.Getenv("GEE_PAY_CLIENT_SECRET"),
		GrantType:    "client_credentials",
	}
	url := os.Getenv("GEE_PAY_AUTH_URL")
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	log.Printf("Body: %s", string(body))
	log.Printf("URL: %s", os.Getenv("GEE_PAY_AUTH_URL"))
	response, statusCode, err := services.SendRequest("POST", url, nil, body)
	if err != nil {
		return nil, err
	}

	if statusCode != 200 {
		return nil, fmt.Errorf("request failed with status code: %d", statusCode)
	}

	var authResponse GeePayAuthResponse
	err = json.Unmarshal(response, &authResponse)
	if err != nil {
		return nil, err
	}

	return &authResponse, nil
}

func GeePayCollection(data GeePayCollectionRequest) (*CollectionResponse, error) {
	// Marshal request body
	transactionRef, err := services.GenerateUUIDv4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate transaction reference: %w", err)
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	token, err := GeePayGenerateToken()

	if err != nil {
		log.Println("Error generating token: %w", err)
		return nil, err
	}

	// Prepare headers
	headers := map[string]string{
		HeaderContentType:    "application/json",
		HeaderAccept:         "application/json",
		HeaderClientID:       os.Getenv("GEE_PAY_CLIENT_ID"),
		HeaderAuthorization:  "Bearer " + token.AccessToken,
		HeaderTransactionRef: transactionRef,
		HeaderCallbackURL:    os.Getenv("GEE_PAY_CALLBACK_URL"),
	}

	// Send request
	response, statusCode, err := services.SendRequest("POST", os.Getenv("GEE_PAY_COLLECTION_REQUEST_URL"), headers, body)
	log.Printf("Response: %s", string(response))
	if err != nil {
		log.Println("Error sending request: %w", err)
		return nil, err
	}

	log.Println("Status code:", statusCode)

	// Check for HTTP status code
	if statusCode != 200 && statusCode != 202 {
		return nil, fmt.Errorf("request failed with status code: %d, response: %s", statusCode, string(response))
	}

	// Parse response
	var collectionResponse CollectionResponse
	if err := json.Unmarshal(response, &collectionResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &collectionResponse, nil
}

func GeePayTransactionEnquiry(transactionRef string) (*TransactionStatusResponse, error) {
	baseURL := os.Getenv("GEE_PAY_TRANSACTION_ENQUIRY_URL")
	url := fmt.Sprintf("%s/%s", baseURL, transactionRef)

	token, err := GeePayGenerateToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	headers := map[string]string{
		HeaderContentType:   "application/json",
		HeaderAccept:        "application/json",
		HeaderAuthorization: "Bearer " + token.AccessToken,
	}

	body, err := json.Marshal(headers)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	response, statusCode, err := services.SendRequest("GET", url, headers, body)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	log.Println("Status code:", statusCode)
	if statusCode != 200 {
		return nil, fmt.Errorf("request failed with status code: %d, response: %s", statusCode, string(response))
	}

	var transactionResponse TransactionStatusResponse
	if err := json.Unmarshal(response, &transactionResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &transactionResponse, nil
}
