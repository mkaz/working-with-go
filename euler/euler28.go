/**
  *
  * Problem 28
  *
  * Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:
  *
  *   21* 22  23  24  25*
  *   20   7*  8   9* 10
  *   19   6   1*  2  11
  *   18   5*  4   3* 12
  *   17* 16  15  14  13*
  *
  * It can be verified that the sum of the numbers on the diagonals is 101.
  *
  * What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
  *

    Pattern for 5 x 5 is first row
    n = 5
    5 x 5 = 25 (max)

    Sum(n) = (25-4*0) + (25-4*1) + (25-4*2) + 25-(4*3)
    Sum(n) = (n*n-((n-1)*0) + (n*n-((n-1)*1) + (n*n-((n-1)*2)) + n*n-((n-1)*3)

    Sum(n-2) = (3*3-2*0) + (3*3-2*1) + (3*3-2*2) + (3*3-2*3)

    Sum(n-2*2) = 1


  Answer: 669171001

  **/

package main

import "fmt"
import "time"

func SumSpiral(num int) int {
	var c = 0
	for i := 0; i < 4; i++ { // once around the corners
		c += num*num - (num-1)*i
	}
	return c
}

func main() {
	var start = time.UnixNano()
	fmt.Printf("Starting Euler #28 \n")

	var sum int = 0
	var x = 1001

	for ; x > 1; x -= 2 {
		sum += SumSpiral(x)
	}

	sum += 1 // dont forget the center

	var end = time.UnixNano()
	var duration = (end - start) / 1000000
	fmt.Printf("Sum: %d \n", sum)
	fmt.Printf("Calculated in %vms \n", duration)

}
