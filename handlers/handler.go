package handlers

import (
	"forum/models"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

var db *gorm.DB
var store *sessions.CookieStore

func SetDB(database *gorm.DB) {
	db = database
}

func SetStore(s *sessions.CookieStore) {
	store = s
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = "./templates/" + tmpl + ".html"
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func PageIndex(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	user, ok := session.Values["user"]
	data := map[string]interface{}{
		"User": user,
	}
	if !ok {
		data["User"] = ""
	}

	var posts []models.Post
	db.Preload("User").Preload("Comments.User").Find(&posts)
	data["Posts"] = posts

	renderTemplate(w, "index", data)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		user := models.User{
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
		renderTemplate(w, "register", nil)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		var user models.User
		email := r.FormValue("email")
		password := r.FormValue("password")
		result := db.Where("email = ? AND password = ?", email, password).First(&user)
		if result.Error != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		session, _ := store.Get(r, "session")
		session.Values["user"] = user.Username
		session.Values["userID"] = user.ID
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		renderTemplate(w, "login", nil)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "user")
	delete(session.Values, "userID")
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userID, ok := session.Values["userID"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		post := models.Post{
			Title:   r.FormValue("title"),
			Content: r.FormValue("content"),
			UserID:  userID.(uint),
		}
		result := db.Create(&post)
		if result.Error != nil {
			http.Error(w, "Unable to create post", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		renderTemplate(w, "create_post", nil)
	}
}

func ViewPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var post models.Post
	if err := db.Preload("User").Preload("Comments.User").First(&post, vars["id"]).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	session, _ := store.Get(r, "session")
	user, ok := session.Values["user"]
	data := map[string]interface{}{
		"Post": post,
		"User": user,
	}
	if !ok {
		data["User"] = ""
	}
	renderTemplate(w, "view_post", data)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userID, ok := session.Values["userID"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		vars := mux.Vars(r)
		postIDStr := vars["id"]
		postID, err := strconv.ParseUint(postIDStr, 10, 32)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}
		r.ParseForm()
		comment := models.Comment{
			Content: r.FormValue("content"),
			PostID:  uint(postID),
			UserID:  userID.(uint),
		}
		result := db.Create(&comment)
		if result.Error != nil {
			http.Error(w, "Unable to create comment", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/post/"+postIDStr, http.StatusSeeOther)
	}
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userID, ok := session.Values["userID"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := db.First(&post, postID).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	var like models.Like
	result := db.Where("user_id = ? AND post_id = ?", userID, postID).First(&like)
	if result.Error == nil {
		db.Delete(&like)
		post.Likes--
	} else {
		like = models.Like{UserID: userID.(uint), PostID: &post.ID}
		db.Create(&like)
		post.Likes++
	}

	if err := db.Save(&post).Error; err != nil {
		http.Error(w, "Unable to update post like", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/post/"+vars["id"], http.StatusSeeOther)
}

func LikeComment(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userID, ok := session.Values["userID"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	vars := mux.Vars(r)
	commentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	var comment models.Comment
	if err := db.First(&comment, commentID).Error; err != nil {
		http.Error(w, "Comment not found", http.StatusNotFound)
		return
	}

	var like models.Like
	result := db.Where("user_id = ? AND comment_id = ?", userID, commentID).First(&like)
	if result.Error == nil {
		db.Delete(&like)
		comment.Likes--
	} else {
		like = models.Like{UserID: userID.(uint), CommentID: &comment.ID}
		db.Create(&like)
		comment.Likes++
	}

	if err := db.Save(&comment).Error; err != nil {
		http.Error(w, "Unable to update comment like", http.StatusInternalServerError)
		return
	}

	postID := vars["postID"]
	http.Redirect(w, r, "/post/"+postID, http.StatusSeeOther)
}
