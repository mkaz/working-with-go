---
title: Control Structures
layout: page
order: 5
---

# Control Structures

Using If, Else, Switch and Conditionals in Go.


## If, Else

Go is not picky, conditionals don't require parentheses.

```go
if num > 3 {
	fmt.Println("Many")
}
```

Go is picky, `else` must be on the same line as closing `if` bracket.

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

Using switch and case conditionals in Go. Note: Go auto breaks each case, this avoids the typical issue of cascading when forgetting a break statement.

```go
switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Many")
	}
}
```

An alternative switch syntax allows you to check beyond equality:

```go
switch {
	case num == 1:
		fmt.Println("One")
	case num == 2:
		fmt.Println("Two")
    case num > 2:
		fmt.Println("Many")
	default:
        fmt.Println("Less than 1")
	}
}
```
