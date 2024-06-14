package main

import (
	"fmt"
	"forum/handlers"
	"forum/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func main() {
	// Initialisation de la base de données
	db, err := gorm.Open(sqlite.Open("forum.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Échec de la connexion à la base de données :", err)
		return
	}

	// Migrer les modèles
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Like{}, &models.Follower{})
	handlers.SetDB(db)
	handlers.SetStore(store)

	// Initialiser le routeur
	r := mux.NewRouter()

	// Routes pour les handlers
	r.HandleFunc("/", handlers.PageIndex).Methods("GET")
	r.HandleFunc("/register", handlers.Register).Methods("GET", "POST")
	r.HandleFunc("/login", handlers.Login).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.Logout).Methods("GET")
	r.HandleFunc("/create-post", handlers.CreatePost).Methods("GET", "POST")
	r.HandleFunc("/post/{id}", handlers.ViewPost).Methods("GET")
	r.HandleFunc("/post/{id}/comment", handlers.CreateComment).Methods("POST")
	r.HandleFunc("/post/{id}/like", handlers.LikePost).Methods("POST")
	r.HandleFunc("/post/{postID}/comment/{id}/like", handlers.LikeComment).Methods("POST")
	r.HandleFunc("/category/{category}", handlers.PostsByCategory).Methods("GET")
	r.HandleFunc("/profile/{username}", handlers.ViewProfile).Methods("GET")
	r.HandleFunc("/profile/{username}/follow", handlers.FollowUser).Methods("POST")
	r.HandleFunc("/profile/{username}/unfollow", handlers.UnfollowUser).Methods("POST")
	r.HandleFunc("/profile/{username}/edit", handlers.EditProfile).Methods("GET", "POST")
	r.HandleFunc("/profile/{username}/followers", handlers.ViewFollowers).Methods("GET")
	r.HandleFunc("/profile/{username}/following", handlers.ViewFollowing).Methods("GET")

	// Servir les fichiers statiques (HTML, CSS, JS)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Démarrer le serveur sur le port 8080
	fmt.Println("Le serveur est en cours d'exécution sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Échec du démarrage du serveur :", err)
	}
}
