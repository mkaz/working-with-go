package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MathFunc func(int, int) int

func main() {
	rand.Seed(time.Now().UnixNano())

	add := func(a int, b int) int { return a + b }
	sub := func(a int, b int) int { return a - b }
	mult := func(a int, b int) int { return a * b }
	fns := []MathFunc{add, sub, mult}

	for i := 0; i < 10; i++ {
		random_index := rand.Intn(3)
		fn := fns[random_index]
		fmt.Printf("%d %d \n", i, fn(i, 1))
	}

}
