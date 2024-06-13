package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint
	User     User
	Comments []Comment
	Likes    int
	Dislikes int
}

type Comment struct {
	gorm.Model
	Content  string
	PostID   uint
	UserID   uint
	User     User
	Likes    int
	Dislikes int
}

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Like struct {
	gorm.Model
	UserID    uint
	PostID    *uint
	CommentID *uint
}
