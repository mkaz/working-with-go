package main

import (
	"fmt"
)

func main() {

	alphas := []string{"abc", "def", "ghi", "jkl"}

	// appending
	alphas = append(alphas, "mno")
	alphas = append(alphas, "pqr", "stu")
	fmt.Printf("%v \n", alphas)

	// length
	fmt.Println("Length: ", len(alphas))

	fmt.Println(alphas[1])

	alpha2 := alphas[1:3]
	fmt.Println(alpha2)

}
