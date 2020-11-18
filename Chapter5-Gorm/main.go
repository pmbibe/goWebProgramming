package main

import "fmt"

func main() {
	connect()

	Db.AutoMigrate(&Post{}, &Comment{})

	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	Db.Create(&post)
	comment := Comment{Content: "Good post!", Author: "Joe"}
	// Db.Create(&comment)
	fmt.Println(post)
	Db.Model(&post).Association("Comments").Append(&comment)
	var readPost Post
	Db.Where(`author = "Sau Sheong"`).First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Joins("left join comments on post_id").Scan(&comments)
	fmt.Println(comments)
}
