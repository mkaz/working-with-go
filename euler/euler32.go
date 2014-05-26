/**
  *
  * Problem 32


We shall say that an n-digit number is pandigital if it makes use of
all the digits 1 to n exactly once; for example, the 5-digit number,
15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 x 186 = 7254, containing
multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity
can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to
only include it once in your sum.

**/

package main

import "fmt"
import "time"

/* import "math"*/
import "strings"

const exp = 5
const max = 10000

func main() {
	t := time.Now()
	answers := NumArray{}

	for x := 1; x < 5000; x++ { // 987 is highest to check
		for y := 1; y < 987; y++ {

			// first get right size numbers
			z := x * y
			str := fmt.Sprintf("%d%d%d", x, y, z)
			if len(str) == 9 {
				keeper := true
				for i := 1; i < 10; i++ {
					numlet := fmt.Sprintf("%d", i)
					if !strings.Contains(str, numlet) {
						keeper = false
					}
				}

				// next get no duplicate numbers
				if keeper {
					answers = answers.append(z)
					fmt.Printf(" %d *  %d =  %d  \n", x, y, z)
				}
			}
		}
	}

	uniqueAnswers := answers.unique()
	fmt.Println("Full Array  : ", answers)
	fmt.Println("Unique Array: ", uniqueAnswers)
	fmt.Println("QED: ", uniqueAnswers.sum())
	fmt.Println("Time  Elapsed: ", time.Since(t))
}

type NumArray []int

func (na NumArray) append(elem int) NumArray { return append(na, elem) }

func (na NumArray) exists(elem int) bool {
	for _, el := range na {
		if elem == el {
			return true
		}
	}
	return false
}

func (na NumArray) elemCount(elem int) (c int) {
	for _, el := range na {
		if elem == el {
			c += 1
		}
	}
	return c
}

func (na NumArray) unique() (a NumArray) {
	for _, el := range na {
		if !a.exists(el) {
			a = append(a, el)
		}
	}
	return a
}

func (na NumArray) sum() (sum int) {
	for _, el := range na {
		sum += el
	}
	return sum
}
