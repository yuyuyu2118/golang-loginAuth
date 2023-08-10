package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Age      int
	Username string
	Password string
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	connStr := "host=localhost port=60185 user=postgres password=ryuusei0618 dbname=users sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var loginRequest LoginRequest
			err := json.NewDecoder(r.Body).Decode(&loginRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			var u User
			row := db.QueryRow("SELECT * FROM users WHERE username=$1 AND password=$2",
				loginRequest.Username, loginRequest.Password)
			err = row.Scan(&u.Id, &u.Name, &u.Email, &u.Age, &u.Username, &u.Password)
			if err == sql.ErrNoRows {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]bool{"success": false})
				return
			} else if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]bool{"success": true})
		}
	})

	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":60180", handler))
}
