---
title: Control Structures
layout: post
date: 2014-05-05
---

# Control Structures

Using If, Else, Switch and Conditionals in Go.


## If, Else

Go is not picky, conditionals don't require parentheses

```go
if num > 3 {
	fmt.Println("Many")
}
```

Go is picky, `else` must be on the same line as closing `if` bracket

```go
if num == 1 {
	fmt.Println("One")
} else if num == 2 {
	fmt.Println("Two")
} else {
	fmt.Println("Many")
}
```

## Switch Statement

Using switch and case conditionals in Go. Note, Go will auto break from different cases, there is not a cascade issue.

```go
switch {
	case num == 1:
		fmt.Println("One")
	case num == 2:
		fmt.Println("Two")
	case num > 2:
		fmt.Println("Many")
	default:
		fmt.Println("Thrown over boat")
	}
}
```

