package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./data/my_database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	InitDB(db)

	router := mux.NewRouter()
	router.HandleFunc("/register", RegisterHandler).Methods("POST")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/post", CreatePostHandler).Methods("POST")
	router.HandleFunc("/comment", CreateCommentHandler).Methods("POST")
	router.HandleFunc("/posts", ListPostsHandler).Methods("GET")

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}
