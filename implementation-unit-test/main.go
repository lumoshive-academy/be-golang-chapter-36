// main.go
package main

import (
    "encoding/json"
    "net/http"
)

const (
    username = "admin"
    password = "password"
)

type Response struct {
    Message string `json:"message"`
}

func BasicAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user, pass, ok := r.BasicAuth()
        if !ok || user != username || pass != password {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
    var res Response
    if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

func main() {
    mux := http.NewServeMux()
    mux.Handle("/post", BasicAuthMiddleware(http.HandlerFunc(PostHandler)))
    http.ListenAndServe(":8080", mux)
}