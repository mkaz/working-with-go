/**
 * read-dir-recursive.go
 *
 * Read directory and recursively print on all files and directories
 * See: http://golang.org/pkg/path/filepath/
 *
 */
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// main function
func main() {
	path := "./"
	// Pass function as a parameter is OK in Go.
	// Walk takes ( path, function ) and goes through path
	// and calls function for each entry
	filepath.Walk(path, Walker)
}

// walker function that gets called per file, receives filename
// fileinfo object and an error if one occurred
func Walker(fn string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Walker Error: ", err)
		return nil
	}

	if fi.IsDir() {
		fmt.Println("Directory: ", fn)
	} else {
		fmt.Println("File: ", fn)
	}
	return nil
}
