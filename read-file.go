package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	filename := "../extras/rabbits.txt"

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file", filename)
	}

	// Note: content returned as []bytes not string
	//       must cast to string first
	fmt.Println(string(content))

}
