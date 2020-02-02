---
title: Interfaces
layout: page
order: 24
---

# Interfaces

An interface in Go is a type that defines a set of methods. Interfaces are the core of Go's type system. A type is said to satisfy an interface if it implements all the methods defined in that interface.


## Empty Interface

The simplest example is the empty `interface{}`. This empty interface defines no methods thus all types implement it. This is often used to define a function that accepts any type, for example `fmt.Print` accepts any type and defined using

```go
func Print(a ...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...)
}
```

The `...` is the spread operator, it allows for any number of `interface{}` types as arguments to the function.


## Interface with Methods

Here is an example interface defining a single method. This defines a `Vehicle` interface as one that has the method `Alert()`

```go
type Vehicle interface {
	Alert() string
}
```

Any type that implements the `Alert()` method is said to satisfy the `Vehicle` interface.

When defining a struct you do not explicitly specify the interface it implements. Go determines automatically if a given type satisfies an interface if it defines all the methods.

```go
type Car struct { }

func (c Car) Alert() string {
	return "Honk! Honk!"
}
```

Defining another type:

```go
type Bicycle struct { }

func (b Bicycle) Alert() string {
	return "Ring! Ring!"
}
```

With the above definitions, you can create an array of Vehicles like so:

```go
package main

import "fmt"

type Vehicle interface {
	Alert() string
}

type Car struct { }

func (c Car) Alert() string {
	return "Honk! Honk!"
}

type Bicycle struct { }

func (b Bicycle) Alert() string {
	return "Ring! Ring!"
}

func main() {
	c := Car{}
	b := Bicycle{}

	vehicles := []Vehicle{c, b}
	for _, v := range vehicles {
		fmt.Println(v.Alert())
	}
}
```
