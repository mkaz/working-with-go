---
title: Guessing Game
layout: page
order: 18
---

# Guessing Game

An example of the guessing game picking a number between 1-100.


## Random umber

The [`math/rand`](https://golang.org/pkg/math/rand/) standard library is used to generate the random number.

Note, you must seed the generator, for random this is easiest done using `time.Now().UnixNano()`. You can also seed the generator with a known value, such as a user id to get a deterministic random number.

```go
rand.Seed(time.Now().UnixNano())
fmt.Println(rand.Intn(100))
```

## Example Game

Here is the example guessing game

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// automatically called on start
func init() {
	// new random seed
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// variable to store guess
	var guess int

	// variable to store number of guesses
	var count int

	// pick number to guess, add 1 since Intn gives [0,99]
	num := rand.Intn(100) + 1

	fmt.Println(">> I'm thinking of a number between 1-100 << ")

	// guessing loop
	for {
		// prompt user for guess
		fmt.Print("Guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err == nil {
			count += 1 // increment guess counter
			if guess > num {
				fmt.Println(" >> Too high ")
			} else if guess < num {
				fmt.Println(" >> Too low ")
			} else {
				fmt.Printf("Correct! It took you %d guesses!\n", count)
				break
			}
		} else { // an error with input
			fmt.Println(">> Please input a number")
		}
	}
}
```

