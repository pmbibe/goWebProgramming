package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post6 struct {
	XMLName  xml.Name  `xml:"post"`
	ID       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}
type Comment struct {
	ID      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func testRule6() {
	xmlFile, err := os.Open("Rule6.xml")
	if err != nil {
		fmt.Println("Error opening XML file: ", err)
		return
	}
	defer xmlFile.Close()
	// xmlData, err := ioutil.ReadAll(xmlFile)
	// if err != nil {
	// fmt.Println("Error reading XML data: ", err)
	// return
	// }
	// var post Post6
	// Unmarshal for small data
	// xml.Unmarshal(xmlData, &post)
	// Decoder
	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
				fmt.Println(comment)
			}
		}
	}
	// fmt.Println(post)
}
