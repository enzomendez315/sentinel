package server

import "net/http"

func Run() error {
	mux := http.NewServeMux()
	return http.ListenAndServe(":8080", mux)
}
