package services

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func SendRequest(method, url string, headers map[string]string, body []byte) ([]byte, int, error) {
	log.Println("Sending request to: %s", url)
	log.Println("Method: %s", method)
	log.Println("Headers: %v", headers)
	log.Println("Body: %v", string(body))
	// Create request with optional body
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error creating request: %w", err)
		return nil, 0, fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers if provided
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Create a client with timeout
	client := &http.Client{Timeout: 15 * time.Second}

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error creating request: %w", err)
		return nil, 0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	log.Println("Response:", resp)

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: %w", err)
		return nil, resp.StatusCode, fmt.Errorf("failed to read response body: %w", err)
	}

	return respBody, resp.StatusCode, nil
}
