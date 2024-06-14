package handlers

import (
	"fmt"
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
			Title:    r.FormValue("title"),
			Content:  r.FormValue("content"),
			Category: r.FormValue("category"),
			UserID:   userID.(uint),
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

func PostsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]

	var posts []models.Post
	if err := db.Preload("User").Preload("Comments.User").Where("category = ?", category).Find(&posts).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	session, _ := store.Get(r, "session")
	user, ok := session.Values["user"]
	data := map[string]interface{}{
		"Posts":    posts,
		"User":     user,
		"Category": category,
	}
	if !ok {
		data["User"] = ""
	}

	renderTemplate(w, "category_posts", data)
}

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	var posts []models.Post
	db.Where("user_id = ?", user.ID).Find(&posts)

	var followers []models.Follower
	var following []models.Follower
	db.Preload("Follower").Where("follows_id = ?", user.ID).Find(&followers)
	db.Preload("Follows").Where("follower_id = ?", user.ID).Find(&following)

	data := map[string]interface{}{
		"ProfileUser":    user,
		"Posts":          posts,
		"Followers":      followers,
		"Following":      following,
		"FollowersCount": len(followers),
		"FollowingCount": len(following),
	}

	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Unable to get session", http.StatusInternalServerError)
		return
	}
	currentUser, ok := session.Values["user"]
	currentUserID := session.Values["userID"]
	if ok {
		data["CurrentUser"] = currentUser
		data["CurrentUserID"] = currentUserID
	} else {
		data["CurrentUser"] = ""
		data["CurrentUserID"] = uint(0)
	}

	// Check if the current user is following the profile user
	var follower models.Follower
	if db.Where("follower_id = ? AND follows_id = ?", currentUserID, user.ID).First(&follower).Error == nil {
		data["IsFollowing"] = true
	} else {
		data["IsFollowing"] = false
	}

	renderTemplate(w, "profile", data)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Unable to get session", http.StatusInternalServerError)
		return
	}
	currentUserID, ok := session.Values["userID"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]

	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	// Vérifiez si l'utilisateur suit déjà
	var follower models.Follower
	if err := db.Where("follower_id = ? AND follows_id = ?", currentUserID, user.ID).First(&follower).Error; err == nil {
		// Déjà suivi, pas besoin de suivre à nouveau
		http.Redirect(w, r, "/profile/"+username, http.StatusSeeOther)
		return
	}

	follower = models.Follower{
		FollowerID: currentUserID.(uint),
		FollowsID:  user.ID,
	}

	if err := db.Create(&follower).Error; err != nil {
		http.Error(w, "Unable to follow user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/profile/"+username, http.StatusSeeOther)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Unable to get session", http.StatusInternalServerError)
		return
	}
	currentUserID, ok := session.Values["userID"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]

	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	// Supprimer la relation de suivi
	if err := db.Where("follower_id = ? AND follows_id = ?", currentUserID, user.ID).Delete(&models.Follower{}).Error; err != nil {
		http.Error(w, "Unable to unfollow user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/profile/"+username, http.StatusSeeOther)
}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Unable to get session", http.StatusInternalServerError)
		return
	}
	currentUserID, ok := session.Values["userID"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]

	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	if user.ID != currentUserID.(uint) {
		http.Error(w, "You do not have permission to edit this profile", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		user.Username = r.FormValue("username")
		user.Email = r.FormValue("email")
		password := r.FormValue("password")

		if password != "" {
			user.Password = password // In a real application, hash the password before storing it
		}

		if err := db.Save(&user).Error; err != nil {
			http.Error(w, "Unable to update profile", http.StatusInternalServerError)
			return
		}

		// Mettre à jour la session avec le nouveau nom d'utilisateur
		session.Values["user"] = user.Username
		session.Save(r, w)

		http.Redirect(w, r, fmt.Sprintf("/profile/%s", user.Username), http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"User": user,
	}

	renderTemplate(w, "edit_profile", data)
}

func ViewFollowers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	var followers []models.Follower
	db.Preload("Follower").Where("follows_id = ?", user.ID).Find(&followers)

	data := map[string]interface{}{
		"ProfileUser": user,
		"Followers":   followers,
	}

	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Unable to get session", http.StatusInternalServerError)
		return
	}
	currentUser, ok := session.Values["user"]
	currentUserID := session.Values["userID"]
	if ok {
		data["CurrentUser"] = currentUser
		data["CurrentUserID"] = currentUserID
	} else {
		data["CurrentUser"] = ""
		data["CurrentUserID"] = uint(0)
	}

	renderTemplate(w, "followers", data)
}

func ViewFollowing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	var following []models.Follower
	db.Preload("Follows").Where("follower_id = ?", user.ID).Find(&following)

	data := map[string]interface{}{
		"ProfileUser": user,
		"Following":   following,
	}

	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Unable to get session", http.StatusInternalServerError)
		return
	}
	currentUser, ok := session.Values["user"]
	currentUserID := session.Values["userID"]
	if ok {
		data["CurrentUser"] = currentUser
		data["CurrentUserID"] = currentUserID
	} else {
		data["CurrentUser"] = ""
		data["CurrentUserID"] = uint(0)
	}

	renderTemplate(w, "following", data)
}
