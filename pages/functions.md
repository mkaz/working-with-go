---
title: Functions
layout: page
order: 4
---

# Functions

Examples on how to use functions in Go.

## Basic Function

An example of a basic function accepting one parameter and no return value. The parameter type must be specified in the function definition.

```go
func Echo(s string) {
	fmt.Println(s)
}
```

## Function with Return Value

Defining a function with a return value, you must specify the type for the return value.

```go
func Say(s string) string {
	phrase := "Hello " + s
	return phrase
}
```

## Named Returned Value

You can define a function with a named return variable. By using a named variable it initializes the variables. Also, by using a named variable you do not need to include the variable in return statement it will return the current value of the variable on return.

```go
func Say(s string) (phrase string) {
	phrase = "Hello " + s
	return
}
```

## Multiple Parameters

Function with multiple parameters and return values

```go
func Divide(x, y float64) (float64, float64) {
	q := math.Trunc(x / y)
	r := math.Mod(x, y)
	return q, r
}
```

Function with multiple parameters and named return values. If the types are the same you can specify the type once at the end

```go
func Divide2(x, y float64) (q, r float64) {
	q = math.Trunc(x / y)
	r = math.Mod(x, y)
	return
}
```

## Variadic Parameters

A variadic function is a function that accepts an arbitrary number of arguments. Here is an example function accepting any number of ints. A slice is created out of the parameters passed in.

```go
func Sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
```

You can call the function with multiple parameters:

```go
sum := Sum(1,3,5,7)
fmt.Println(sum)
```

You can also call using the spread operator:

```go
nums := []int{1, 2, 3, 4, 5 }
sum := Sum(nums...)
fmt.Println(sum)
```

