---
title: Structs
layout: page
order: 7
---

# Working with Structs

A struct is a data type in Go that is a collection of fields. A struct is the closest Go gets to objects.

## Define Struct

```go
// define Dog struct
type Dog struct {
	Name  string
	Color string
}
```

## Create Instance

Create an instance of a struct type

```go
Spot := Dog{Name: "Spot", Color: "brown"}
```

An alternative way, create a blank instance, each field is set to its zero value. Access struct fields using dot notation for getting or setting.

```go
var Rover Dog
Rover.Name = "Rover"
Rover.Color = "light tan with dark patches"
```

