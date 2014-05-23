package main

import "fmt"

func main() {

	alphas := []string{"abc", "def", "ghi"}

	// standard for loop
	for i := 0; i < len(alphas); i += 1 {
		fmt.Printf("%d: %s \n", i, alphas[i])
	}

	// using range, same as above
	for i, val := range alphas {
		fmt.Printf("%d: %s \n", i, val)
	}

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	//
	// for {
	//   fmt.Print(".")
	// }

}
