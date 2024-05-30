package main

import (
	"github.com/gorilla/mux"
)

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/register", registerHandler).Methods("GET", "POST")
	r.HandleFunc("/login", loginHandler).Methods("GET", "POST")
	r.HandleFunc("/create-post", createPostHandler).Methods("GET", "POST")

	// Ajoutez d'autres routes selon les besoins
}
