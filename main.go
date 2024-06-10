package main

import (
	"fmt"
	"forum/handlers"
	"net/http"
)



func main() {


    http.HandleFunc("/", handlers.Page_index)


   // Démarrer le serveur sur le port 8080
   fmt.Println("Le serveur est en cours d'exécution sur http://localhost:8080")
   if err := http.ListenAndServe(":8080", nil); err != nil {
       fmt.Println("Échec du démarrage du serveur :", err)
   }
}
