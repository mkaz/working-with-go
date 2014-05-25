/**
  * command-line-args.go
  * A program showing how to use command-line arguments and flags
  */

package main
import (
	"flag"
	"fmt"
	"os"
)

var conf string
var limit int

func main() {

	// command-line args are stored as slice in os.Args
	// first argument in list is program itself
	num_args := len( os.Args )
	
	// flag package provides parsing of command-line parameters
	flag.StringVar(&conf, "conf", "default value", "text description")
	flag.IntVar(&limit, "limit", 5, "text description")
	flag.Parse()

	fmt.Println("Conf: ", conf)
	fmt.Println("Limit: ", limit)

	// last command-line argument
	fmt.Println("Last Item: ", os.Args[num_args-1])
}

// $ go run command-line-args.go
// >> Conf: default value
// >> Conf: 5

// $ go run command-line-args.go --conf=Foo --limit=8 filename
// >> Conf: Foo
// >> Conf: 8
