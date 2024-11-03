// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Buat request HTTP palsu
	req := httptest.NewRequest("GET", "http://example.com/", nil)

	// Buat recorder untuk menangkap response
	w := httptest.NewRecorder()

	// Panggil handler dengan request dan recorder
	HelloHandler(w, req)

	// Periksa status code response
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
	}

	// Periksa body response
	expectedBody := "Hello, World!\n"
	body := w.Body.String()
	if body != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, body)
	}
}

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
