/**
 * functions.go
 * Examples on how to use functions in go
 */

// remember all programs need a package
package main

// import packages
import (
	"fmt"  // standard output formatting
	"math" // example uses math package
)

// Basic function accepting one parameter
// You must specify the type of variables expected
// The format is function_name(variable type)
func Echo(s string) {
	fmt.Println(s)
}

// Function with single return value
// the type of the return value is specifed after function declaration
func Say(s string) string {
	phrase := "Hello " + s
	return phrase
}

// Function with single named return value
// You can specify the return variable name, which initializes it
// the := notation is for new variables, and = for initial ones
// Also you do not need to include the variable in return statement
// it will return the current value of the variable at return
func Say2(s string) (phrase string) {
	phrase = "Hello " + s
	return
}

// Function with multiple parameters and return values
func Divide(x, y float64) (float64, float64) {
	q := math.Trunc(x / y)
	r := math.Mod(x, y)
	return q, r
}

// Function with multiple parameters and named return values
// If the types are the same you can specify the type once at the end
func Divide2(x, y float64) (q, r float64) {
	q = math.Trunc(x / y)
	r = math.Mod(x, y)
	return
}

// Main function which will just call the above defined functions
func main() {

	// call the basic Echo function
	Echo("Bonjour tout le monde")

	// call the Say function which returns a string
	fmt.Println(Say("Gopher"))

	// test function with multiple return value
	q, r := Divide(11, 3)

	// this example uses Printf to format output, %v can be used for any type
	fmt.Printf("Quotient: %v, Remainder %v \n", q, r)

}
