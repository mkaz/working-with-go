/**
 * working-with-slices.go
 *
 * Go has an array type, but most interactions are with slices which are built
 * off arrays. I don't worry too much about the name and just use them
 * The specification for slices are "[]T" where T is the type.
 * So "[]string" is a set of strings, "[]int" is a set of integers, and so on.
 */

package main

import "fmt"

func main() {

	// initializes an empty array
	var empty []int

	// initialize array with values
	alphas := []string{"abc", "def", "ghi", "jkl"}

	// slices can not be modified, a new copy needs to be made
	// here is a example that appends elements to a slice
	empty = append(empty, 123)
	empty = append(empty, 456)
	fmt.Printf("%v \n", empty)

	// append multiple values
	alphas = append(alphas, "pqr", "stu")
	fmt.Printf("%v \n", alphas)

	// get length of a slice
	fmt.Println("Length: ", len(alphas))

	// retrieve a single element from slice
	fmt.Println(alphas[1])

	// retrieve a slice of a slice
	alpha2 := alphas[1:3]
	fmt.Println(alpha2)

	// check if element exists in array
	// there is no function to determine this
	// the only method is to iterate over the array
	// see elemExists func defined below
	if elemExists("def", alphas) {
		fmt.Println("Exists!")
	}

}

func elemExists(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
