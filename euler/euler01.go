/**
 *
 * Problem 1
 *
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
 * The sum of these multiples is 23.

 * Find the sum of all the multiples of 3 or 5 below 1000.
 *

 * Answer: 233168

 **/
package main

import "fmt"

const MAX = 1000

func main() {

	sum := 0

	for i := 1; i < MAX; i++ {
		sum += check(i)
	}

	fmt.Printf("Total: %d \n", sum)

}

func check(num int) int {
	if (num%3 == 0) || (num%5 == 0) {
		return num
	}

	return 0
}
