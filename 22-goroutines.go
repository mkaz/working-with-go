// Originally from: https://gobyexample.com/channels
// [Reprinted under license](https://creativecommons.org/licenses/by/3.0/).
// Poetic changes were made.

package main

import (
	"fmt"
	"time"
)

func printSlowly(s string, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i, s)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {

	// This is a normal function call.
	// Main() will finish this off before continuing.
	printSlowly("directly functioning", 3)

	// The go functions below will each spin off to happen each in their own thread,
	// Meaning they'll be called _concurrently_.

	// Calling the named function as a go routine.
	go printSlowly("red fish goroutine", 3)
	go printSlowly("blue fish goroutine", 3)

	// Call an anonymous function as a go routine.
	go func(ss string, nn int) {
		for i := 0; i < nn; i++ {
			fmt.Println(i, ss)
			time.Sleep(150 * time.Millisecond)
		}
	}("anony fish goroutine", 3)

	// Waits for a button to be pushed.
	// Try commenting this!
	var input string
	fmt.Scanln(&input) // Just push RETURN to finish the program.
	fmt.Println("DONE.")
}
