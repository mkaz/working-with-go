---
title: Strings
parent: 1465
template: page-tut.php
order: 2
---

# Working with Strings

A set of examples using Golang and creating and manipulating strings. The [`strings`](http://golang.org/pkg/strings/) standard library contains most of the functions you'll use to work with string.

You can create and set a variable in one step using `:=` operator. Go will automatically determine the type by the assignment.

```go
str := "This is an example string"
```

In general, Go does not operate on standard values using an object oriented dot notation like JavaScript and Python. Instead, the function typically involve passing in the variable you are working on.

```go
exists := strings.Contains(str, "example")
```


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

Strings in Go are an array of characters and can referenced as such.

```go
str := "abcdefghijklmnopqrstuvwxyz"
fmt.Println("Chars 3-10: " + str[3:10])

// printing out first 5 characters
fmt.Println("First Five: " + str[:5])

// printing out from 13 to end
fmt.Println("From 13 on: " + str[13:])
```

The Go Playground is a useful online environment you can run Go code in your browser. See the above example in the Go Playground here: [https://play.golang.org/p/DC7R7XKzZ5G](https://play.golang.org/p/DC7R7XKzZ5G)

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

## String Replace

You can replace a string using `strings.Replace`

```go
str := "The blue whale loves blue fish."
newstr := strings.Replace(str, "blue", "red", 1)
>> "The red whale loves blue fish."
```

The `strings.Replace` function requires passing in how many replacements it should do. You can pass in `-1` to replace all.

```go
str := "The blue whale loves blue fish."
newstr := strings.Replace(str, "blue", "red", -1)
>> "The red whale loves red fish."
```

As of, Go v1.12, there is a new `strings.ReplaceAll` that does not have the fourth argument, it will replace all.


## Strings HasPrefix and HasSuffix

The Go functions for checking if a strings starts with, or ends with, another string are `strings.HasPrefix()` and `strings.HasSuffix()`.

```go
path := "/home/mkaz"
isAbsolute := strings.HasPrefix(path, "/")

dir := "/home/mkaz/"
hasTrailingSlash := strings.HasSuffix(path, "/")
```

