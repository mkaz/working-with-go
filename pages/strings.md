---
title: Strings
layout: page
order: 2
---

# Working with Strings

A set of examples working with strings in Golang. The `strings` standard library contains most of the functions you'll use to work with string. <div class="sidenote">See [strings package](http://golang.org/pkg/strings/) documentation</div>

Create and set a variable in one line using `:=` operator. Go will automatically determine the type by the assignment.

```go
str := "This is an example string"
```

In general, Go does not use a standard object oriented notation to operate on values. This differs from languages like JavaScript and Python. Instead in Go, you typically pass in the variable you are working on to the function. Here is an example using the `Contains` function from the `strings` package:

```go
exists := strings.Contains(str, "example")
```
Here is a full example program, that uses the `ToLower` function to convert a string to lowercase, and `Contains` to check if a string contains another.

```go
package main

import (
	"fmt"     // for standard output
	"strings" // for manipulating strings
)

func main() {

	// create a string variable
	str := "HI, I'M UPPER CASE"

	// convert to lower case
	lower := strings.ToLower(str)

	// output to show its really lower case
	fmt.Println(lower)

	// check if string contains another string
	if strings.Contains(lower, "case") {
		fmt.Println("Yes, exists!")
	}
}
```

## Strings as Array of Characters

Strings in Go are an array of characters and can be referenced as such.

```go
str := "abcdefghijklmnopqrstuvwxyz"
fmt.Println("Chars 3-10: " + str[3:10])

// printing out first 5 characters
fmt.Println("First Five: " + str[:5])

// printing out from 13 to end
fmt.Println("From 13 on: " + str[13:])
```

<span class="tip">💡</span> The Go Playground is a useful online environment that runs Go code in your browser. See the above example in the Go Playground here: [https://play.golang.org/p/DC7R7XKzZ5G](https://play.golang.org/p/DC7R7XKzZ5G). I often use the playground to test out a bit of logic I'm working on.

## Split Strings

Split a string on a specific character or word

```go
sentence := "I'm a sentence made up of words"
words := strings.Split(sentence, " ")
fmt.Printf("%v \n", words)
```

If you were splitting on whitespace, using Fields is better because it will split on all whitespace characters, not just a space.

```go
sentence := "I'm a sentence made up of words"
fields := strings.Fields(sentence)
fmt.Printf("%v \n", fields)
```

## Join an Array of Strings

In an OOP world, the `Join` function would be defined in an array package. In Go, however, it is part of the strings package, you pass in an array of strings and the separator you want to join them.

```go
lines := []string{ "one", "two", "three" }
str := string.Join(lines, ",")
fmt.Println(str)
>> one, two, three
```

## String Replace

To replace a string with another use `strings.Replace`

```go
str := "The blue whale loves blue fish."
newstr := strings.Replace(str, "blue", "red", 1)
>> "The red whale loves blue fish."
```

The `strings.Replace` function requires passing in how many replacements it should do. Pass in `-1` to replace all occurrences.

```go
str := "The blue whale loves blue fish."
newstr := strings.Replace(str, "blue", "red", -1)
>> "The red whale loves red fish."
```

Go version 1.12 introduced the `strings.ReplaceAll` function that does not have the fourth argument -- it, as it says, replaces all occurrences.


## Strings HasPrefix and HasSuffix

The Go functions for checking if a strings starts with, or ends with, another string are `strings.HasPrefix()` and `strings.HasSuffix()`.

```go
path := "/home/mkaz"
isAbsolute := strings.HasPrefix(path, "/")

dir := "/home/mkaz/"
hasTrailingSlash := strings.HasSuffix(path, "/")
```

