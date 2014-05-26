/**
  *
  * Problem 30

  Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

  1634 = 1^4 + 6^4 + 3^4 + 4^4
  8208 = 8^4 + 2^4 + 0^4 + 8^4
  9474 = 9^4 + 4^4 + 7^4 + 4^4
  As 1 = 1^4 is not a sum it is not included.

  The sum of these numbers is 1634 + 8208 + 9474 = 19316.

  Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
**/

package main

import "fmt"
import "time"
import "math"

const exp = 5
const max = 10000

func main() {
	t := time.Now()
	qed := 0
	var start = t.UnixNano()
	fmt.Printf("Starting Euler #30 \n")

	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					for e := 0; e <= 9; e++ {
						for f := 0; f <= 9; f++ {

							num := a*100000 + b*10000 + c*1000 + d*100 + e*10 + f
							a2 := int(math.Pow(float64(a), float64(exp)))
							b2 := int(math.Pow(float64(b), float64(exp)))
							c2 := int(math.Pow(float64(c), float64(exp)))
							d2 := int(math.Pow(float64(d), float64(exp)))
							e2 := int(math.Pow(float64(e), float64(exp)))
							f2 := int(math.Pow(float64(f), float64(exp)))
							num2 := a2 + b2 + c2 + d2 + e2 + f2
							if num == num2 {
								fmt.Printf("Found: %d \n", num)
								qed += num
							}
						}
					}
				}
			}
		}
	}

	qed = qed - 1 // by defn 1 is not applicable

	tt := time.Now()
	var end = tt.UnixNano()
	duration := (end - start) / 1000000
	fmt.Printf("QED: %d \n", qed)
	fmt.Printf("Calculated in %vms \n", duration)

}
