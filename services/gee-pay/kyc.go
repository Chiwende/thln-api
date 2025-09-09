package services

import (
	"encoding/json"
	"log"
	"os"

	services "qinsights.com/thln/services"
)

type KYCSuccessResponse struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Status      string `json:"status"`
		Provider    string `json:"provider"`
		PhoneNumber string `json:"phone_number"`
		Names       string `json:"names"`
	} `json:"data"`
}

type KYCErrorResponse struct {
	Code   string `json:"code"`
	Errors struct {
		PhoneNumber []string `json:"phone_number"`
	} `json:"errors"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func KYCService(phone string) (*KYCSuccessResponse, *KYCErrorResponse, error) {
	url := os.Getenv("GEE_PAY_KYC_URL") + phone
	log.Println("URL" + url)
	token, err := GeePayGenerateToken()
	if err != nil {
		log.Println("Error generating token: %w", err)
		return nil, nil, err
	}
	headers := map[string]string{
		"Authorization": "Bearer " + token.AccessToken,
		"Content-Type":  "application/json",
		"Accept":        "application/json",
	}
	response, statusCode, err := services.SendRequest("GET", url, headers, nil)
	if err != nil {
		log.Println("Error sending request: %w", err)
		return nil, nil, err
	}
	if statusCode != 200 {
		var errorResponse KYCErrorResponse
		err = json.Unmarshal(response, &errorResponse)
		if err != nil {
			return nil, nil, err
		}
		return nil, &errorResponse, nil
	}
	var successResponse KYCSuccessResponse
	err = json.Unmarshal(response, &successResponse)
	if err != nil {
		return nil, nil, err
	}
	return &successResponse, nil, nil
}
