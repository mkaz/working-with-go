/*

  Project Euler #41
  Marcus Kazmierczak, marcus@mkaz.com

  We shall say that an n-digit number is pandigital if it makes use
  of all the digits 1 to n exactly once. For example, 2143 is a 4-digit
  pandigital and is also prime.

  What is the largest n-digit pandigital prime that exists?

*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var primes []int64

func main() {

	t := time.Now()
	fmt.Printf("Starting Euler #41 \n")

	//const MAX = 987654321
	const MAX = 999999999999999

	buildPrimes(MAX) // create all prime candidates

	for _, prime := range primes {
		numstr := strconv.FormatInt(prime, 10)
		if isPanDigital(numstr) {
			fmt.Printf("Found: %v \n", prime)
		}
	}

	fmt.Printf("--\nTime Elapsed: %v \n", time.Since(t))

}

func isPanDigital(s string) bool {
	l := len(s)
	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	check := nums[:l]
	for _, c := range check {
		if !strings.Contains(s, c) {
			return false
		}
	}
	return true
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

func buildPrimes(num int64) {
	stop := int(math.Floor(math.Sqrt(float64(num))))
	primes = append(primes, 2) // seed primes
	for i := 3; i < stop; i += 2 {
		isPrime(int64(i)) // builds array
	}
}
