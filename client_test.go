package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockLoadConfig() Config {
	return Config{
		MockyURL: "https://example.com",
	}
}

func TestSendMessages(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqBody map[string]string
		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"status":"success"}`))
	}))
	defer mockServer.Close()

	config := MockLoadConfig()
	config.MockyURL = mockServer.URL

	messages := &[]Message{
		{Recipient: "user@example.com", Content: "Hello!"},
	}

	status, err := SendMessages(messages)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if status != http.StatusAccepted {
		t.Errorf("Expected status 202, got %d", status)
	}
}
