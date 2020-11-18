package main

import "fmt"

func readData() {
	db := connect()
	defer db.Close()

	slect, _ := db.Query("SELECT * FROM posts")
	defer slect.Close()
	for slect.Next() {
		var post Post
		err := slect.Scan(&post.ID, &post.Content, &post.Author)
		errHandler(err)
		fmt.Printf("%d %s %s \n", post.ID, post.Content, post.Author)

	}

}
