// main_test.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostHandler(t *testing.T) {
	// JSON body untuk request
	requestBody := Response{Message: "Hello, World!"}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Buat request HTTP palsu dengan basic auth
	req := httptest.NewRequest("POST", "http://example.com/post", bytes.NewReader(body))
	req.SetBasicAuth("admin", "password")

	// Buat recorder untuk menangkap response
	w := httptest.NewRecorder()

	// Panggil handler dengan request dan recorder
	handler := BasicAuthMiddleware(http.HandlerFunc(PostHandler))
	handler.ServeHTTP(w, req)

	// Periksa status code response
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	// Periksa body response
	var responseBody Response
	if err := json.NewDecoder(w.Body).Decode(&responseBody); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if responseBody.Message != requestBody.Message {
		t.Errorf("expected body message %q, got %q", requestBody.Message, responseBody.Message)
	}
}

func TestPostHandlerUnauthorized(t *testing.T) {
	// JSON body untuk request
	requestBody := Response{Message: "Hello, World!"}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Buat request HTTP palsu tanpa basic auth
	req := httptest.NewRequest("POST", "http://example.com/post", bytes.NewReader(body))

	// Buat recorder untuk menangkap response
	w := httptest.NewRecorder()

	// Panggil handler dengan request dan recorder
	handler := BasicAuthMiddleware(http.HandlerFunc(PostHandler))
	handler.ServeHTTP(w, req)

	// Periksa status code response
	res := w.Result()
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("expected status code %d, got %d", http.StatusUnauthorized, res.StatusCode)
	}
}
