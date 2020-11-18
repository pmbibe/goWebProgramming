package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

func storeG(data interface{}, filename string) {
	//Create byte.Buffer struct to use for Read and Write method
	buffer := new(bytes.Buffer)
	// Using gob.NewEncoder to used Encode function
	encoder := gob.NewEncoder(buffer)
	//Encode data -> buffer (binary)
	err := encoder.Encode(data)
	errHandler(err)
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	errHandler(err)
}
func loadG(data interface{}, filename string) {
	//Load data from binary file
	raw, err := ioutil.ReadFile(filename)
	errHandler(err)
	//Create Buffer from data
	buffer := bytes.NewBuffer(raw)
	//Create new Decoder by NewDecoder
	dec := gob.NewDecoder(buffer)
	//Decode buffer (binary) -> speccialy struct
	err = dec.Decode(data)
	errHandler(err)
}
func RAWFromGob() {
	post := Post{ID: 1, Content: "Hello World!", Author: "Sau Sheong"}
	//Saving a post to binary file (post1 is binary file)
	storeG(post, "post1")
	var postRead Post
	//Loading data from binary file to Post struct
	loadG(&postRead, "post1")
	fmt.Println(postRead)
}
