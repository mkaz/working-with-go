/**
 *
 * Problem 10
 *
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 * Find the sum of all the primes below two million.
 *
 * Answer: 142913828922
 **/

package main

import "fmt"
import "time"

var primes []int64

func isPrime(num int64) bool {
	var stop = len(primes)
	for i := 0; i < stop; i++ {
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

func main() {
	primes = append(primes, 2)
	var start = time.Nanoseconds()
	fmt.Printf("Starting Euler #10 \n")

	const MAX = 2000000
	var sum int64 = 2 // start at include 2
	var i int64 = 3   // start at 3

	for ; i < MAX; i += 2 { // count by 2's
		if isPrime(i) {
			sum += i
		}
	}

	var end = time.Nanoseconds()
	var duration = (end - start) / 1000000
	fmt.Printf("Sum: %d \n", sum)
	fmt.Printf("Calculated in %vms \n", duration)

}
