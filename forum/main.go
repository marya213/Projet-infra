package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))
var dbPath = os.Getenv("DB_PATH")

func main() {
    db := initDB(dbPath)
    defer db.Close()

    r := mux.NewRouter()
    registerRoutes(r)

    log.Println("Starting server on :8080")
    http.ListenAndServe(":8080", r)
}
