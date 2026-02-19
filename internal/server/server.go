package server

import (
	"fmt"
	"net/http"
)

func Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			fmt.Printf("Error writing response: %v", err)
		}
	})

	return http.ListenAndServe(":8080", mux)
}
