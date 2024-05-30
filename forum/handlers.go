package main

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer et valider les données du formulaire
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Hasher le mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}

	// Enregistrer l'utilisateur dans la base de données
	_, err = db.Exec("INSERT INTO users (email, username, password_hash) VALUES (?, ?, ?)", email, username, hashedPassword)
	if err != nil {
		http.Error(w, "Email ou nom d'utilisateur déjà pris", http.StatusConflict)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer et valider les données du formulaire
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Récupérer l'utilisateur de la base de données
	var hashedPassword string
	err := db.QueryRow("SELECT password_hash FROM users WHERE email = ?", email).Scan(&hashedPassword)
	if err != nil {
		http.Error(w, "Email ou mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	// Vérifier le mot de passe
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		http.Error(w, "Email ou mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	// Créer une nouvelle session
	session, _ := store.Get(r, "session-name")
	session.Values["email"] = email
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	email, ok := session.Values["email"].(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Récupérer les données du formulaire
	title := r.FormValue("title")
	content := r.FormValue("content")
	categories := r.Form["categories"] // C'est un tableau de catégories sélectionnées

	// Récupérer l'ID de l'utilisateur
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE email = ?", email).Scan(&userID)
	if err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusInternalServerError)
		return
	}

	// Insérer le post dans la base de données
	result, err := db.Exec("INSERT INTO posts (user_id, title, content) VALUES (?, ?, ?)", userID, title, content)
	if err != nil {
		http.Error(w, "Erreur lors de la création du post", http.StatusInternalServerError)
		return
	}

	// Récupérer l'ID du post créé
	postID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}

	// Associer les catégories au post
	for _, categoryID := range categories {
		_, err = db.Exec("INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)", postID, categoryID)
		if err != nil {
			http.Error(w, "Erreur lors de l'association des catégories", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
