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
	db.AutoMigrate(&models.User{})
	handlers.SetDB(db)

	// Initialiser le routeur
	r := mux.NewRouter()

	// Routes pour les handlers
	r.HandleFunc("/", handlers.PageIndex).Methods("GET")
	r.HandleFunc("/register", handlers.Register).Methods("GET", "POST")
	r.HandleFunc("/login", handlers.Login).Methods("GET", "POST")

	// Servir les fichiers statiques (HTML, CSS, JS)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./templates/"))))

	// Démarrer le serveur sur le port 8080
	fmt.Println("Le serveur est en cours d'exécution sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Échec du démarrage du serveur :", err)
	}
}
