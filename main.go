package main

import (
	"encoding/json"
	"net/http"

	"github.com/rs/cors"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = map[string]string{
	"test": "test",
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	password, ok := users[req.Username]
	if !ok || password != req.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Logged in successfully"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handleLogin)

	handler := cors.Default().Handler(mux) // CORS middlewareを追加

	http.ListenAndServe(":60180", handler)
}
