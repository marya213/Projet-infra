package main

import (
	"database/sql"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Post struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

type Comment struct {
	ID      int    `json:"id"`
	PostID  int    `json:"post_id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

// CreateUser inserts a new user into the database.
func CreateUser(db *sql.DB, email, username, password string) error {
	_, err := db.Exec("INSERT INTO users(email, username, password) VALUES(?, ?, ?)", email, username, password)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByEmail retrieves a user from the database by email.
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	row := db.QueryRow("SELECT id, email, username, password FROM users WHERE email = ?", email)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreatePost inserts a new post into the database.
func CreatePost(db *sql.DB, userID int, title, content, category string) error {
	_, err := db.Exec("INSERT INTO posts(user_id, title, content, category) VALUES(?, ?, ?, ?)", userID, title, content, category)
	if err != nil {
		return err
	}
	return nil
}

// GetAllPosts retrieves all posts from the database.
func GetAllPosts(db *sql.DB) ([]Post, error) {
	rows, err := db.Query("SELECT id, user_id, title, content, category FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Category)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// CreateComment inserts a new comment into the database.
func CreateComment(db *sql.DB, postID, userID int, content string) error {
	_, err := db.Exec("INSERT INTO comments(post_id, user_id, content) VALUES(?, ?, ?)", postID, userID, content)
	if err != nil {
		return err
	}
	return nil
}

// GetCommentsByPostID retrieves all comments for a given post ID from the database.
func GetCommentsByPostID(db *sql.DB, postID int) ([]Comment, error) {
	rows, err := db.Query("SELECT id, post_id, user_id, content FROM comments WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

// InitDB initializes the database with the required tables.
func InitDB(db *sql.DB) {
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT UNIQUE,
        username TEXT,
        password TEXT
    );
    CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        title TEXT,
        content TEXT,
        category TEXT,
        FOREIGN KEY(user_id) REFERENCES users(id)
    );
    CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        post_id INTEGER,
        user_id INTEGER,
        content TEXT,
        FOREIGN KEY(post_id) REFERENCES posts(id),
        FOREIGN KEY(user_id) REFERENCES users(id)
    );
    `
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
}
