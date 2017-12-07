/**
 * read-write-files.go
 *
 * Examples reading and writing to a file using io/ioutil
 * See: http://golang.org/pkg/io/ioutil/
 *
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	filename := "./extras/rabbits.txt"

	// read in file, one command loads file into content variable
	// if an error occurs returns it as the second return value
	// if no error, err will be nil
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		// log is a nifty little utility which can also output
		// in this case, a fatal log will halt the program
		log.Fatalln("Error reading file", filename)
	}

	// content returned as []byte not string
	// so must cast to string first and then can display
	fmt.Println(string(content))

	// One of the above errors can be if the file does not exist
	// You can check explicitly if file exists without trying to
	// open using
	if _, err := os.Stat("junk.foo"); os.IsNotExist(err) {
		fmt.Println(">>>")
		fmt.Println("File: junk.foo does not exist")
	}

	// write back to new file
	// see documentation for which methods take what type
	outfile := "output.txt"
	err = ioutil.WriteFile(outfile, content, 0644)
	if err != nil {
		log.Fatalln("Error writing file: ", err)
	} else {
		fmt.Println(">>>")
		fmt.Println("Created: output.txt")
	}

	// write to an existing file appending the content
	af, err := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error appending to file:", err)
	}
	defer af.Close()
	if _, err = af.WriteString("Appending this text"); err != nil {
		log.Fatalln("Error writing to file:", err)
	}

}
