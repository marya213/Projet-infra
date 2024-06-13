package main

import (
	"fmt"
	"forum/handlers"
	"forum/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialisation de la base de données
	db, err := gorm.Open(sqlite.Open("forum.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Échec de la connexion à la base de données :", err)
		return
	}

	// Migrer les modèles
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	handlers.SetDB(db)

	// Initialiser le routeur
	r := mux.NewRouter()

	// Routes pour les handlers
	r.HandleFunc("/", handlers.Page_index).Methods("GET")
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/posts/{id:[0-9]+}", handlers.GetPostByID).Methods("GET")
	r.HandleFunc("/posts/category/{category}", handlers.GetPostsByCategory).Methods("GET")
	r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/comments", handlers.CreateComment).Methods("POST")
	r.HandleFunc("/posts/{id:[0-9]+}/like", handlers.LikePost).Methods("POST")
	r.HandleFunc("/posts/{id:[0-9]+}/dislike", handlers.DislikePost).Methods("POST")
	r.HandleFunc("/users", handlers.ListUsers).Methods("GET") // Nouvelle route pour lister les utilisateurs

	// Servir les fichiers statiques (HTML, CSS, JS)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./template/"))))

	// Démarrer le serveur sur le port 8080
	fmt.Println("Le serveur est en cours d'exécution sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Échec du démarrage du serveur :", err)
	}
}
