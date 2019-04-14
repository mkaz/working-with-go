---
title: Functions
layout: post
date: 2014-05-04
---

# Functions

Examples on how to use functions in go

## Basic Function

An example of a basic function accepting one parameter and no return value. The parameter type must be specified in the function defintion.

```go
func Echo(s string) {
	fmt.Println(s)
}
```

## Function with Return Value

Definining a function with a return value, you must specify the type for the return value.

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

