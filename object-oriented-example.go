package main

import (
	"fmt"
)

// create dog type
type Dog struct {
	Name  string
	Color string
}

// create method for dog
func (d Dog) Call() {
	fmt.Printf("Come here %s dog, %s \n", d.Color, d.Name)
}

func main() {
	Spot := Dog{Name: "Spot", Color: "brown"}
	Spot.Call()
}
