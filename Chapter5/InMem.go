package main

import "fmt"

func InMem() {
	PostByID = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)
	post1 := Post{ID: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{ID: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{ID: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{ID: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}
	store(post1)
	store(post2)
	store(post3)
	store(post4)
	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
