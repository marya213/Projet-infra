package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const port = ":8080"

// Artist représente un artiste
type Artist struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}

// HomeHandler est le gestionnaire pour l'endpoint de la page d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// CSSHandler est le gestionnaire pour le fichier CSS
func CSSHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "style.css")
}

// ArtistsHandler est le gestionnaire pour l'API des artistes
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	// Effectuer une requête GET vers l'API des artistes
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "Erreur lors de la requête GET: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Vérifier le code de statut de la réponse
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Statut de réponse non OK: "+resp.Status, http.StatusInternalServerError)
		return
	}

	// Décode la réponse JSON en une slice d'artistes
	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		http.Error(w, "Erreur lors du décodage JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Définir l'en-tête de contenu JSON
	w.Header().Set("Content-Type", "application/json")

	// Encoder les données des artistes en JSON et les écrire dans la réponse
	err = json.NewEncoder(w).Encode(artists)
	if err != nil {
		http.Error(w, "Erreur lors de l'encodage JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	fmt.Println("(http://localhost:8080) - Serveur Start sur le port", port)

	// Gestionnaire pour le fichier CSS
	http.HandleFunc("/style.css", CSSHandler)

	// Gestionnaire pour l'image d'arrière-plan
	http.Handle("/background.jpg", http.FileServer(http.Dir(".")))

	// Gestionnaire pour l'API des artistes
	http.HandleFunc("/artists", ArtistsHandler)

	// Gestionnaire pour l'endpoint de la page d'accueil
	http.HandleFunc("/", HomeHandler)

	// Démarrer le serveur
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur de démarrage du serveur:", err)
	}
}
