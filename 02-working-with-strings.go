/**
 * working-with-strings.go
 *
 * Working with strings, manipulating, creating etc
 * Most strings functions are stored in standard library "strings"
 * See: http://golang.org/pkg/strings/
 */

// standard main package
package main

// Note: if you include a package but don't use it, the Go compiler will barf
import (
	"fmt"     // for standard output
	"strings" // for manipulating strings
)

// this function gets run on program execution
func main() {

	// create a string variable
	str := "HI, I'M UPPER CASE"

	// convert to lower case
	lower := strings.ToLower(str)

	// output to show its really lower case
	fmt.Println(lower)

	// check if string contains another string
	if strings.Contains(lower, "case") {
		fmt.Println("Yes, exists!")
	}

	// split a string on a specific character or word
	sentence := "I'm a sentence made up of words"
	words := strings.Split(sentence, " ")
	fmt.Printf("%v \n", words)

	// If you were splitting on whitespace, using Fields is better because
	// it will split on more than just the space, but all whitespace chars
	fields := strings.Fields(sentence)
	fmt.Printf("%v \n", fields)

}

// run program in your terminal using
// $ go run working-with-strings.go
