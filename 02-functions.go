package main

import (
	"fmt"
	"math"
)

// function with single return value
func Say(s string) (phrase string) {
	if s == "" {
		return
	}
	phrase = "Hello " + s
	return
}

// function with multiple return values
func Divide(x, y float64) (q, r float64) {
	q = math.Trunc(x / y)
	r = math.Mod(x, y)
	return q, r
}

func main() {
	// test function
	fmt.Println(Say(""))
	fmt.Println(Say("Gopher"))

	// test function with multiple return value
	q, r := Divide(11, 3)
	fmt.Printf("Quotient: %v, Remainder %v \n", q, r)
}
