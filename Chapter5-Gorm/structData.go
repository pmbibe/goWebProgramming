package main

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        int
	Content   string
	Author    string `sql:"not null`
	Comments  []Comment
	CreatedAt time.Time
}
type Comment struct {
	ID      int
	Content string
	Author  string `sql:"not null"`
	PostID  int    `sql:"index"`
}

var Db *gorm.DB
