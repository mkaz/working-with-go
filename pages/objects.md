---
title: Objects
layout: page
order: 11
---

# Object Oriented Programming

Go does not have objects and classes, so is not a true object oriented programming. However, you can add methods to types. This gives you most of the same ability without the added complexity of full OOP.

A method is a function with a _receiver_ defined first. Using the `Dog` type from the [structs example](/structs.html).

```go
package main

import "fmt"

type Dog struct {
	Name  string
	Color string
}

func (d Dog) Call() {
	fmt.Printf("Here comes a %s dog, %s.\n", d.Color, d.Name)
}

func main() {

	// create instance of dog
	Spot := Dog{Name: "Spot", Color: "brown"}

	// call object method
	Spot.Call()
}
```


## Limitations

You can define methods on custom types, but are limited to types that are defined in the same package as the method. So you cannot define a method on built-in types, you would first need to extend.

```go
package main

import "fmt"

type MyInt int

func (i MyInt) Double() MyInt {
	return i * 2
}

func main() {
	x := MyInt(3)
	fmt.Println(x.Double())
}
```

