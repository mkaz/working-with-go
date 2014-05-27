/**
 * guess-game.go
 *
 * Guess a number between 1-100
 *
 */

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
