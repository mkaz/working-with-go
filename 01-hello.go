/**
 * Welcome to Go Code!
 *
 * A set of go programs to help you learn golang by reading code
 * Kinda like the immersion foreign language classes but better
 * because parlez vous golang es muy bueno
 *
 * hello.go
 * A minimal first program
 */

// all programs must be part of a package and main is the default package
package main

// Import standard library packages, they are not included by default.
// The standard library is your best friend in Go, it can do just about
// everything. See all the glory -> http://golang.org/pkg/

// fmt library provides basic standard out formatting functions
import "fmt"

// the main package requires a main() function which gets called when run
func main() {

	// define a variable, all variables have types, but the go compiler will
	// automatically detect the type if possible, this is a string
	phrase := "Hola Mundo!"

	// call the Println method from fmt package, which writes to standard out
	fmt.Println(phrase)
}

// To run this go program in your terminal, requires having go installed
// For installation info see: http://golang.org/doc/install
// once installed run the following, from the terminal ($ is command-prompt)
// $ go run 01-hello.go
