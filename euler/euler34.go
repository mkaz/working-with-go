/**
 *
 * Problem 34
 *
 *  145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
 *
 *  Find the sum of all numbers which are equal to the sum of the factorial of their digits.
 *
 *  Note: as 1! = 1 and 2! = 2 are not sums they are not included.
 *
 */

package main

import (
	"fmt"
	"strconv"
	"time"
)

const max = 10000000

func factorial(num int) (qed int) {
	qed = 1
	for i := 1; i <= num; i++ {
		qed *= i
	}
	return qed
}

func main() {
	t := time.Now()
	total := 0

	for c := 0; c <= max; c++ {
		s := strconv.Itoa(c)
		sum := 0
		for i := 0; i < len(s); i++ {
			j, _ := strconv.Atoi(string(s[i]))
			sum += factorial(j)
		}

		if (sum == c) && (sum != 1) && (sum != 2) {
			total += c
			fmt.Printf("Num: %d String: %s ", c, s)
			fmt.Printf("\n")
		}

	}

	fmt.Println("------------------")
	fmt.Println("QED: ", total)
	fmt.Println("Time Elapsed: ", time.Since(t))
}
