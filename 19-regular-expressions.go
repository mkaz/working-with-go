/**
 * regular-expressions.go
 *
 * A set of examples using regular expressions in Go.
 * See package documentation: http://golang.org/pkg/regexp/
 */
package main

import (
	"fmt"
	"regexp"
)

func basic_regexes() {
	// create regular expression pattern
	// pattern, match 1 or more numbers
	pattern := "[0-9]+"

	// test string
	str := "The 12 monkeys ate 48 bananas"

	// compile pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	// 1. Test compiled pattern matches string
	if re.MatchString(str) {
		fmt.Println("Yes, matched a number")
	} else {
		fmt.Println("No, no match")
	}

	// 2. Return first match
	result_two := re.FindString(str)
	fmt.Println("Number matched:", result_two)

	// 3. Return n matches, use -1 to find all matches
	results_three := re.FindAllString(str, 2)
	for i, v := range results_three {
		fmt.Printf("Match %d: %s\n", i, v)
	}

	// 4. Replace matches
	results_four := re.ReplaceAllString(str, "xx")
	fmt.Println("Result:", results_four)
}

func case_insensitive() {
	// the format for flags is a bit different than typical regex
	// in golang the flag precedes the pattern, the syntax is not great
	// here is a case insensitive flag
	ptn := `(?i)^t.`
	str := "To be or not to be"

	re, err := regexp.Compile(ptn)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	// match string
	result := re.FindString(str)
	fmt.Println("Result:", result)
}


func sub_matches() {
	// Submatches are what Go calls the matches that
	// are grabbed by (.*) inside of a regular expression

	str1 := "Hello @world@ Match"
	sub_re, _ := regexp.Compile("@(.*)@")
	m := sub_re.FindStringSubmatch(str1)
	// FindStringSubmatch returns [@world@ world]
	// so to just get the submatch you would use
	if len(m) > 1 {
		fmt.Println(m[1]) // submatch
	}

	// Escaping Special Characters
	// if you wanted to match brackets or other special
	// characters and try to just escape them like so
	// "\[(.*)\]" you will get error
	// unknown escape sequence: [

	// you need to double up the slashes or a nicer
	// solution is to wrap in ticks instead of quotes
	str2 := "A [word] here and [there]"
	esc_pattern := `\[(.*?)\]`
	esc_re, _ := regexp.Compile(esc_pattern)

	// this will only find the first
	fmt.Println(esc_re.FindStringSubmatch(str2))

	// use FindAll with second parameter for # of matches -1 = all
	fmt.Println(esc_re.FindAllStringSubmatch(str2, -1))
}

func main() {
	basic_regexes()
	case_insensitive()
	sub_matches()
}
