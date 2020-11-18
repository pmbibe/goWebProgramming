package main

type Post struct {
	ID      int
	Content string
	Author  string
}

//PostByID One ID - One Post
var PostByID map[int]*Post

//PostsByAuthor One Author can have many Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostByID[post.ID] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)

}
