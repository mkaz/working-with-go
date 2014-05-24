/**
  * object-oriented-example.go
  * Create basic object and method
  */

package main
import "fmt"

// define Dog object type
type Dog struct {
	Name  string
	Color string
}

// define method for Dog object
func (d Dog) Call() {
	fmt.Printf("Come here %s dog, %s \n", d.Color, d.Name)
}

func main() {

	// create object and set properties
	Spot := Dog{Name: "Spot", Color: "brown"}

	// call object method
	Spot.Call()

}
