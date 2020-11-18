package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func CSVFile() {
	//Write file CSV
	csvFile, err := os.Create("Posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	allPosts := []Post{
		Post{ID: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{ID: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{ID: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{ID: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.ID), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}

	}
	writer.Flush()
	//Read file CSV
	csvRead, err := os.Open("Posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvRead.Close()
	reader := csv.NewReader(csvRead)
	// Set the FieldsPerRecord field in the reader to be a negative number, which indicates that
	// you aren’t that bothered if you don’t have all the fields in the record. If FieldsPerRecord
	// is a positive number, then that’s the number of fields you expect from each
	// record and Go will throw an error if you get less from the CSV file. If FieldsPerRecord
	// is 0, you’ll use the number of fields in the first record as the FieldsPerRecord value
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	errHandler(err)
	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{ID: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].ID)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)

}
