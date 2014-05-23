package main

import (
	"io/ioutil"
	"log"
)

func main() {

	filename := "output.txt"
	content := "text to write to the file"
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatalln("Error writing file: ", err)
	}

}
