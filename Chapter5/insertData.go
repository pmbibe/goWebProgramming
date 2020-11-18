package main

import "fmt"

func insertData() {
	db = connect()
	defer db.Close()
	var posts []Post
	post1 := Post{ID: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{ID: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{ID: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{ID: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}
	posts = append(posts, post1)
	posts = append(posts, post2)
	posts = append(posts, post3)
	posts = append(posts, post4)
	for i := range posts {
		query := fmt.Sprintf(`INSERT INTO posts VALUES (%d,"%s","%s")`, posts[i].ID, posts[i].Content, posts[i].Author)
		fmt.Println(query)
		db.Query(query)
	}
}
