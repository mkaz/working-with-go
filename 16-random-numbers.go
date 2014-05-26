/**
 * random-numbers.go
 *
 * Examples generating random numbers with Go
 *
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MathFunc func(int, int) int

func main() {

	// random package has a method called Intn(n) which will return
	// a random number between [0, n)
	// However, if you don't seed the generator, I always get 81, 87, 47
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println("---")

	// an easy way to seed is use nano time which will vary enough
	// this is useful to create deterministic random numbers, for example
	// based of a user_id number
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println("---")

	// Here are a few other random number methods Go provides
	// As always see the package docs http://golang.org/pkg/math/rand/
	fmt.Println("Random Int:", rand.Int())
	fmt.Println("Random Float:", rand.Float32())
	fmt.Println("Random Permutation of N ints:", rand.Perm(8))

	// A not so random function in the math/rand package
	// The NormFloat64 method returns a number based on Normal distrib
	// mean=0, stddev=1 -- if called enough and plotted will give bell curve
	fmt.Println("Normalized:", rand.NormFloat64())
}
