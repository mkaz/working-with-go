/*

  Project Euler #42
  Marcus Kazmierczak, marcus@mkaz.com


	The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1);
	so the first ten triangle numbers are:

	1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

	By converting each letter in a word to a number corresponding to its alphabetical position
	and adding these values we form a word value.

	For example, the word value for SKY is 19 + 11 + 25 = 55 = t10.
	If the word value is a triangle number then we shall call the word a triangle word.

	Using words.txt a text file containing nearly two-thousand common English words,
	how many are triangle words?


*/

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var numbers []int

func main() {

	t := time.Now()
	fmt.Printf("Starting Euler #42 \n")

	getTriangleNumbers() // create all prime candidates

	// read words.txt file
	bytes, _ := ioutil.ReadFile("./data/42-words.txt")
	words := strings.Split(string(bytes), ",")

	total := 0
	triangle := 0
	for _, word := range words {
		word = strings.Replace(word, "\"", "", -1)
		word = strings.ToLower(word)
		num := convertWordToNum(word)

		if exists(num, numbers) {
			triangle += 1
		}

		total += 1

	}
	// loop through each word convert to number
	// check if number = triangle

	fmt.Println("--")
	fmt.Println("Total Words   : ", total)
	fmt.Println("Total Triangle: ", triangle)
	fmt.Println("--")
	fmt.Printf("Time Elapsed: %v \n", time.Since(t))

}

func exists(num int, a []int) bool {
	for _, v := range a {
		if num == v {
			return true
		}
	}
	return false
}

func convertWordToNum(w string) int {
	bytes := []byte(w)
	sum := 0
	for _, b := range bytes {
		n := int(b) - 96
		sum += n
	}
	return sum
}

func getTriangleNumbers() {

	for i := 1; i <= 1000; i++ {
		tn := 0.5 * float32(i) * (float32(i) + 1.0)
		numbers = append(numbers, int(tn))
	}
}
