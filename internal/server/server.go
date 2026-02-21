package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/enzomendez315/sentinel/internal/endpoint"
	"github.com/google/uuid"
)

func Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			fmt.Printf("Error writing response: %v", err)
		}
	})

	mux.HandleFunc("POST /endpoints", func(w http.ResponseWriter, r *http.Request) {
		var ep endpoint.Endpoint

		// Decode JSON into struct
		err := json.NewDecoder(r.Body).Decode(&ep)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		// Set fields
		ep.ID = uuid.New().String()
		ep.CreatedAt = time.Now()
		ep.UpdatedAt = time.Now()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		// Encode JSON for response
		err = json.NewEncoder(w).Encode(ep)
		if err != nil {
			fmt.Printf("Error writing response: %v", err)
		}
	})

	fmt.Println("Starting server at http://localhost:8080")

	return http.ListenAndServe(":8080", mux)
}
