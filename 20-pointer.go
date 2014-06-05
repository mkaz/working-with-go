/**
 * pointer.go
 *
 * Use object's pointer to modify object.
 * See: http://tour.golang.org/#28
 */

package main

import "fmt"

// define Dog object type
type Dog struct {
	Name  string
	Color string
}

func main() {

	// create instance of object and set properties
	Spot := Dog{Name: "Spot", Color: "brown"}

	// get pointer of object
	SpotPointer := &Spot

	// modify field through pointer
	SpotPointer.Color = "black"

	fmt.Println(Spot.Color)

}
