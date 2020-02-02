---
title: Numbers
layout: page
order: 2
---

# Working with Numbers

Go is a typed language, that means each variable is given a single specific number type. The common number types in golang are int, float, int64, float64. You need to be aware that each are differnet and function definitions and usage define the precise type expected.

## Convert String to Integer

The most common conversion is a `string` to an `int`, for this you use the `strconv` package. <div class="sidenote">See [strconv package](https://golang.org/pkg/strconv) documentation.</div> The strconv package contains various functions for converting strings to default types.

Here is how to convert a string to an integer in golang:

```go
str := 42
ans, err := strconv.Atoi(str)
if err != nil {
	fmt.Println("Error convert string to integer", err)
}
```

## Convert Integer to a String

Converting an integer to a string is a bit easier because there is no chance of an error. So it is just the single function `strconv.Itoa()`

```go
str := strconv.Itoa(42)
```

## Convert Integer to a Float

The `math` package has numerous helpful functions. <div class="sidenote">See [math package](https://golang.org/pkg/math/) documentation</div> However, almost of the function require using a float64, so conversions are needed.

If you are specifying a number, you can force a float type by using a decimal.
```go
i := 42    // int type
y := 42.0  // float type
```

If an integer a variable, you can cast it to a float using
```go
i := 42
x := float64(i)
```

## Convert a String to a Float

To convert a string to a float, use the `strconv.ParseFloat` function, the second argument is the bitSize, 64 is probably the most common to convert to a float64 type.

```go
pi := "3.14159"
x, err := strconv.ParseFloat(pi, 64)
if err != nil {
	fmt.Println("Error convert string to float", err)
}
```

Note: You can print the type of a variable using `%T` in `fmt.Printf()`

Example printing the type of a variable in golang:
```go
fmt.Printf("Variable x has type %T and value %v", x, x )
```
