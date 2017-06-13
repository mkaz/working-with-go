/**
 * for-loops.go
 * Go uses "for" for all loops, there is no "while" or "do-while" loops
 *
 */

package main

import "fmt"

func main() {

	// creating an array of strings
	alphas := []string{"abc", "def", "ghi"}

	// standard for loop
	for i := 0; i < len(alphas); i++ {
		fmt.Printf("%d: %s \n", i, alphas[i])
	}

	// if iterating over a array, this would be how you would actually do it
	// the standard loop would be used if uncommon step function
	// range returns two values, i -> index, val -> value
	for i, val := range alphas {
		fmt.Printf("%d: %s \n", i, val)
	}

	// if you did not care about the value and only wanted the index
	for i := range alphas {
		fmt.Println(i)
	}

	// if you did not care about the index and only the value, you use the _
	// which means don't save this value
	for _, val := range alphas {
		fmt.Println(val)
	}

	// how to use for like a while
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// infinite loop
	// for {
	//   fmt.Print(".")
	// }

}
