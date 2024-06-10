package models

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email    string `gorm:"unique"`
    Username string `gorm:"unique"`
    Password string
}

type Post struct {
    gorm.Model
    Title     string
    Content   string
    UserID    uint
    User      User
    Comments  []Comment
    Category  string
    Likes     int
    Dislikes  int
}

type Comment struct {
    gorm.Model
    Content string
    PostID  uint
    UserID  uint
    User    User
    Likes   int
    Dislikes int
}
