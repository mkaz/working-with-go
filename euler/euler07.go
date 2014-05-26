/**
 *
 * Problem 7
 *
 * What is the 10,001 prime
 *
 **/

package main

import (
	"fmt"
	"time"
)

var primes []int64

func main() {

	t := time.Now()
	fmt.Printf("Starting Euler #3 \n")

	//const num = 6
	const num = 10001
	const MAX = 1000000

	c := 1 // prime list is seeded with 2

	for n := 3; n < MAX; n += 2 {
		if isPrime(int64(n)) {
			c = c + 1
		}

		if c == num {
			fmt.Printf(" QED: %v \n", n)
			break
		}
	}

	fmt.Printf("--\nTime Elapsed: %v \n", time.Since(t))

}

/* primes is an array of all known previous primes
   use array to check number is prime, if so add to array */
func isPrime(num int64) bool {

	for i := 0; i < len(primes); i++ {
		if num%primes[i] == 0 {
			return false
		}
		// dont need to try more than square root of num
		if primes[i]*primes[i] > num {
			break
		}
	}
	primes = append(primes, num)
	return true
}
