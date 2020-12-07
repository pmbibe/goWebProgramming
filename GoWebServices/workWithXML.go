package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	//Struct tag
	//Key: must be xml - Value: must have double quotes ("")
	//Some rules for the XML struct tags:
	// 1. To store the name of XML element ifself, field named XMLName, type xml.Name
	// 2. To store the attribute of an XML element, define a field
	// with the same name as the XML element tag, and use the
	// struct tag `xml:"<name>,attr"` where name is the name of the XML attribute
	// 3. To store the character data value of the XML element, define a field
	// with the same name as the XML element tag, and use the struct tag:
	// `xml:",chardata"`
	// 4. To get raw XML from XML element, define a field (using any name) and
	// use the struct tag `xml:",innerxml"`
	// 5. If there are no mode flags (like ,attr or ,chardata or ,innerxml) the
	// struct field will be matched with an XML element that has the same name
	// as the struct's name.
	// 6. To get to an XML element directly without specifying the tree structure
	// to get that element, use the struct tag `xml:"a>b>c"`, where a and b are the
	// intermediate elements and c is the node that you want to get to
	XMLName xml.Name `xml:"post"`
	ID      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	XML     string   `xml:",innerxml"`
}
type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func WorkWithXML(file string) {
	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XMl data:", err)
		return
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
