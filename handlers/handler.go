package handlers

import (
	"html/template"
	"net/http"

	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	tmpl = "./templates/" + tmpl + ".html"
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func PageIndex(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		user := User{
			Username: r.FormValue("username"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		result := db.Create(&user)
		if result.Error != nil {
			http.Error(w, "Unable to register user", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		renderTemplate(w, "register")
	}
}

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		var user User
		email := r.FormValue("email")
		password := r.FormValue("password")
		result := db.Where("email = ? AND password = ?", email, password).First(&user)
		if result.Error != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		renderTemplate(w, "login")
	}
}
