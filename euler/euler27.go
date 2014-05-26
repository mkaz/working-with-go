package main

/*
* Project Euler #27
*
* Euler published the remarkable quadratic formula:
*
* n² + n + 41
*
* It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39.
* However, when n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41,
* and certainly when n = 41, 41² + 41 + 41 is clearly divisible by 41.
*
* Using computers, the incredible formula  n² - 79n + 1601 was discovered,
* which produces 80 primes for the consecutive values n = 0 to 79.
* The product of the coefficients, -79 and 1601, is -126479.
*
* Considering quadratics of the form:
*
* n² + an + b, where |a| < 1000 and |b| < 1000
*
* where |n| is the modulus/absolute value of n
* e.g. |11| = 11 and |4| = 4
*
* Find the product of the coefficients, a and b, for the quadratic expression
* that produces the maximum number of primes for consecutive values of n, starting with n = 0.

* Answer:
*  Primes Array: 8483
*  Coeffs: a = -61, b=971 found 71 primes
*  Coeff Product: -59231
*
* Calculation Time:
*  Node: 64s
*  Go  : 31.4s
 */

import "fmt"
import "time"

func main() {

	var start = time.Nanoseconds()

	var max_prime = 80*80 + 80*1000 + 1000
	var primes = PrimeArray(max_prime)
	fmt.Printf("Primes Array: %d \n", len(primes))

	var found_a = 0
	var found_b = 0
	var found_primes = 0

	for a := -1000; a < 1000; a++ {
		// fmt.Printf("A: " + a);
		for b := -1000; b < 1000; b++ {
			var num_primes = 0

			for n := 0; n < 80; n++ {
				var value = n*n + a*n + b
				if exists(value, primes) {
					num_primes += 1
				} else {
					break
				}
			}

			if num_primes > found_primes {
				found_primes = num_primes
				found_a = a
				found_b = b
			}

		}
	}

	var end = time.Nanoseconds()
	var elapsed = (end - start) / 10e8
	fmt.Printf("Elapsed: %vs \n", elapsed)

	// Here's what I found
	var coeff_product = found_a * found_b
	fmt.Printf("Coefficients:  a = %v, b = %v, found %v primes \n", found_a, found_b, found_primes)
	fmt.Printf("Coeff Product: %v \n", coeff_product)

}

func exists(elem int, array []int) bool {
	for _, value := range array {
		if elem == value {
			return true
		}
	}
	return false
}

/* Return an Array of Primes up to max */
func PrimeArray(max int) []int {
	var primes = []int{2, 3, 5}

	for num := 7; num <= max; num += 2 {
		var is_prime = true
		for _, value := range primes {
			if num%value == 0 {
				is_prime = false
				break
			}

			// dont need to try more than square root of num
			if value*value >= num {
				break
			}
		}
		if is_prime {
			primes = append(primes, num)
		}
	}
	return primes
}
