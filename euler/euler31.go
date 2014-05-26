/**
  *
  * Problem 31


In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:

    1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).

It is possible to make £2 in the following way:

    1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p

How many different ways can £2 be made using any number of coins?

For calculating combination for making 10p

1p    2p    5p    10p
10    0     0     0
8     1     0     0
6     2     0     0
4     3     0     0
2     4     0     0
0     5     0     0

5     0     1     0
0     0     2     0
0     0     0     1

3     1     1     0
1     2     1     0


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

	fmt.Println("Time Elapsed: ", time.Since(t))
}
