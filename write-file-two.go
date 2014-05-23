package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalln("Error creating file: ", err)
	}
	// deferring close will close file before program exits
	// good practice to defer right after opening
	defer f.Close()

	for _, str := range []string{"apple", "banana", "carrot"} {
		bytes, err := f.WriteString(str)
		if err != nil {
			log.Fatalln("Error writing string: ", err)
		}
		fmt.Printf("Wrote %d bytes to file\n", bytes)
	}

}
