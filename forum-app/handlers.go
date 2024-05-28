package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	_, err := db.Exec("INSERT INTO users(email, username, password) VALUES(?, ?, ?)", user.Email, user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Implémentation de la gestion de la connexion
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	// Implémentation de la création de post
}

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Implémentation de la création de commentaire
}

func ListPostsHandler(w http.ResponseWriter, r *http.Request) {
	// Implémentation de l'affichage des posts
}
