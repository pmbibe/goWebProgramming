package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func RAW() {
	data := []byte("Hello World! \n")
	// 	Writes to file and
	// reads from file using
	// WriteFile and ReadFile
	err := ioutil.WriteFile("data1", data, 0644)
	errHandler(err)
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))

	//Writes to file and reads from file using
	// the File struct
	file1, _ := os.Create("data2")
	defer file1.Close()
	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file \n", bytes)
	file2, _ := os.Open("data2")
	defer file2.Close()
	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file \n", bytes)
	fmt.Println(string(read2))
}
