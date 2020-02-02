---
title: Slices
layout: page
order: 3
---

# Working with Slices

Go has an array type, but it has a fixed size. Slices in Go are dynamically-sized and more flexible. Slices are the common data type you will work with.

The specification for slices are `[]T` where T is the type, for example `[]string` is a set of strings, or `[]int` is a set of integers.

## Initialize

Initialize an empty slice.

```go
var empty []int
```

Initialize a slice with values.

```go
alphas := []string{"abc", "def", "ghi", "jkl"}
```

## Append

Slices can not be modified, a new copy needs to be made, it is common to copy to the same variable.

```go
var nums []int
nums = append(nums, 203)
nums = append(nums, 302)
fmt.Println(nums)
// [203 302]
```

Append multiple values at once.

```go
alphas := []string{"abc", "def", "ghi", "jkl"}
alphas = append(alphas, "mno", "pqr", "stu")
```

## Common Array Operations

Get length of a slice.

```go
fmt.Println("Length: ", len(alphas))
```

Retrieve a single element from slice

```go
fmt.Println(alphas[1])
```

Retrieve a slice of a slice

```go
alpha2 := alphas[1:3]
fmt.Println(alpha2)
```


## Element Exists in a Slice

There is no built-in function to determine if an element exists in a slice. Here is an example function to check if a string exists in an array of strings.

```go
package main

import "fmt"

func main() {

	alphas := []string{"abc", "def", "ghi", "jkl"}

	if elemExists("def", alphas) {
		fmt.Println("Exists!")
	}
}

func elemExists(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
```

