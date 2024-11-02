// // main_test.go
// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestPostHandler(t *testing.T) {
// 	// JSON body untuk request
// 	requestBody := Response{Message: "Hello, World!"}
// 	body, err := json.Marshal(requestBody)
// 	assert.NoError(t, err, "Failed to marshal request body")

// 	// Buat request HTTP palsu dengan basic auth
// 	req := httptest.NewRequest("POST", "http://example.com/post", bytes.NewReader(body))
// 	req.SetBasicAuth("admin", "password")

// 	// Buat recorder untuk menangkap response
// 	w := httptest.NewRecorder()

// 	// Panggil handler dengan request dan recorder
// 	handler := BasicAuthMiddleware(http.HandlerFunc(PostHandler))
// 	handler.ServeHTTP(w, req)

// 	// Periksa status code response
// 	res := w.Result()
// 	assert.Equal(t, http.StatusOK, res.StatusCode, "expected status code 200 OK")

// 	// Periksa body response
// 	var responseBody Response
// 	err = json.NewDecoder(w.Body).Decode(&responseBody)
// 	assert.NoError(t, err, "Failed to decode response body")
// 	assert.Equal(t, requestBody.Message, responseBody.Message, "expected body message to match request body message")
// }

// func TestPostHandlerUnauthorized(t *testing.T) {
// 	// JSON body untuk request
// 	requestBody := Response{Message: "Hello, World!"}
// 	body, err := json.Marshal(requestBody)
// 	assert.NoError(t, err, "Failed to marshal request body")

// 	// Buat request HTTP palsu tanpa basic auth
// 	req := httptest.NewRequest("POST", "http://example.com/post", bytes.NewReader(body))

// 	// Buat recorder untuk menangkap response
// 	w := httptest.NewRecorder()

// 	// Panggil handler dengan request dan recorder
// 	handler := BasicAuthMiddleware(http.HandlerFunc(PostHandler))
// 	handler.ServeHTTP(w, req)

// 	// Periksa status code response
// 	res := w.Result()
// 	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "expected status code 401 Unauthorized")
// }

package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// Definisikan tabel test dengan berbagai kasus uji
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Positive numbers", 1, 2, 3},
		{"Negative numbers", -1, -1, -2},
		{"Mix of positive and negative", 1, -1, 0},
		{"Zero", 0, 0, 0},
	}

	// Iterasi melalui setiap kasus uji dalam tabel test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
