package main

import (
	"database/sql"
	"errors"
	"fmt"
)

var db *sql.DB

type PostC struct {
	ID       int
	Content  string
	Author   string
	Comments []CommentP
}
type CommentP struct {
	ID      int
	Content string
	Author  string
	Post    *PostC
}

func (comment *CommentP) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}
	insForm, err := db.Prepare("INSERT INTO comments VALUES(?,?,?,?)")
	if err != nil {
		return
	}
	fmt.Println(comment.ID, comment.Content, comment.Author, comment.Post.ID)
	_, err = insForm.Exec(comment.ID, comment.Content, comment.Author, comment.Post.ID)
	return
}
func GetPost(id int) (post PostC, err error) {
	post = PostC{}
	post.Comments = []CommentP{}
	pRow := db.QueryRow("select id, content, author from posts where id = ?", id)
	_ = pRow.Scan(&post.ID, &post.Content, &post.Author)

	rows, err := db.Query("select id, content, author from comments where post_id = ?", id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := CommentP{Post: &post}
		err = rows.Scan(&comment.ID, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	fmt.Println(post.ID, post.Content, post.Author, post.Comments)
	return
}
