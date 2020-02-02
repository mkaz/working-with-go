---
title: For Loops
layout: page
order: 6
---

# Loops

Go uses `for` for all loops, there is no `while` or `do-while` loops.

## For Loop

```go

alphas := []string{"abc", "def", "ghi"}

// standard for loop
for i := 0; i < len(alphas); i++ {
	fmt.Printf("%d: %s \n", i, alphas[i])
}

// counting backwards
for i := len(alphas)-1; i >= 0; i-- {
	fmt.Printf("%d: %s \n", i, alphas[i])
}
```

## Using Range

Iterating over an array is easier using `range` function.

```go
for i, val := range alphas {
	fmt.Printf("%d: %s \n", i, val)
}
```

If you did not care about the value and only wanted the index

```go
for i := range alphas {
	fmt.Println(i)
}
```

If you did not care about the index and only the value, you need to use the `_` character.

```go
for _, val := range alphas {
	fmt.Println(val)
}
```

## Use For like  While

```go
x := 0
for x < 10 {
	fmt.Println(x)
	x++
}
```

## Infinite Loop

```go
for {
	fmt.Print(".")
}
```

