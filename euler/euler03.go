/**
*
* Problem 3
*
* The prime factors of 13195 are 5, 7, 13 and 29.
* What is the largest prime factor of the number 600851475143 ?
*
**/

package main

import (
	"fmt"
	"math"
	"time"
)

var primes []int64

func main() {

	t := time.Now()
	fmt.Printf("Starting Euler #3 \n")

	//const num = 13195
	const num = 600851475143

	buildPrimes(num)
	for _, prime := range primes {
		if num%prime == 0 { // prime factor
			fmt.Printf("Prime Factor: %d \n", prime)
		}
	}

	fmt.Printf("--\nTime Elapsed: %v \n", time.Since(t))

}

func buildPrimes(num int64) {
	stop := int(math.Floor(math.Sqrt(float64(num))))
	primes = append(primes, 2) // seed primes
	for i := 3; i < stop; i += 2 {
		isPrime(int64(i)) // builds array
	}
}

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
