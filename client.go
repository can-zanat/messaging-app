package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SendMessages(messages *[]Message) (int, error) {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	var statusCode int

	for _, msg := range *messages {
		// Prepare the JSON payload
		payload := map[string]string{
			"to":      msg.Recipient,
			"content": msg.Content,
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			return 0, fmt.Errorf("failed to marshal payload: %v", err)
		}

		// Create a new POST request
		req, err := http.NewRequest(
			http.MethodPost,
			config.MockyURL,
			bytes.NewBuffer(jsonData),
		)

		if err != nil {
			return 0, fmt.Errorf("failed to create request: %v", err)
		}

		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			return 0, fmt.Errorf("request error: %v", err)
		}

		// Capture the status code
		statusCode = resp.StatusCode
		if statusCode != http.StatusAccepted {
			return statusCode, fmt.Errorf("received non-200 status code: %d", statusCode)
		}

		resp.Body.Close()
	}

	return statusCode, nil
}
