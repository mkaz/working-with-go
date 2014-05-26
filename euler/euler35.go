/**
 *
 * Problem 35
 *
 * The number, 197, is called a circular prime because all
 * rotations of the digits: 197, 971, and 719, are themselves prime.
 * There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
 * How many circular primes are there below one million?
 *
 */

package main

import (
	"fmt"
	"strconv"
	"time"
)

const max = 1000000

var primes []int

func main() {
	t := time.Now()
	qed := 0
	sum := 0

	// fill prime array
	const MAX = 10000000
	primes = append(primes, 2)
	for i := 3; i < MAX; i += 2 { // count by 2's
		if isPrime(i) {
			sum += 1
		}
	}

	for c := 0; c <= max; c++ {
		s := strconv.Itoa(c)
		allprime := true

		for i := 0; i < len(s); i++ {
			numStr := ""
			for j := 0; j < len(s); j++ {
				ind := j + i
				if (ind) >= len(s) {
					ind = ind - len(s)
				}
				char := string(s[ind])
				numStr += char
			}
			num, _ := strconv.Atoi(numStr)

			if !isPrime(num) {
				allprime = false
				break
			}
		}

		if allprime {
			qed += 1
			fmt.Println(c)
		}

	}

	fmt.Println("------------------")
	fmt.Println("QED: ", qed)
	fmt.Println("Time Elapsed: ", time.Since(t))
}

func isPrime(num int) bool {
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
