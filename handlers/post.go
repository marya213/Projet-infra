package handlers

import (
	"encoding/json"
	"forum/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
    var post models.Post
    json.NewDecoder(r.Body).Decode(&post)

    if err := db.Create(&post).Error; err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(post)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
    var comment models.Comment
    json.NewDecoder(r.Body).Decode(&comment)

    if err := db.Create(&comment).Error; err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(comment)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
    var posts []models.Post
    db.Preload("Comments").Find(&posts)
    json.NewEncoder(w).Encode(posts)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var post models.Post
    if err := db.Preload("Comments").First(&post, id).Error; err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(post)
}

func GetPostsByCategory(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    category := vars["category"]

    var posts []models.Post
    db.Preload("Comments").Where("category = ?", category).Find(&posts)
    json.NewEncoder(w).Encode(posts)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var post models.Post
    if err := db.First(&post, id).Error; err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }

    post.Likes++
    db.Save(&post)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(post)
}

func DislikePost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var post models.Post
    if err := db.First(&post, id).Error; err != nil {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }

    post.Dislikes++
    db.Save(&post)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(post)
}
