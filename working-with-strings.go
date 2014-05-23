package main

import (
	"fmt"
	"strings"
)

func main() {

	// to lower case
	str := "HI, I'M UPPER CASE"
	lower := strings.ToLower(str)
	fmt.Println(lower)

	// contains
	if strings.Contains(lower, "case") {
		fmt.Println("Yes, exists!")
	}

	// split
	sentence := "I'm a sentence made up of words"
	words := strings.Split(sentence, " ")
	fmt.Printf("%v \n", words)

	// fields
	fields := strings.Fields(sentence)
	fmt.Printf("%v \n", fields)

}
