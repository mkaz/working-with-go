package main
import (
    "fmt"
	"os"
    "path/filepath"
)

func main() {
    path := "/etc/"
    filepath.Walk(path, Walker)
}

// walker function that gets called per file
func Walker(fn string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Walker Error: ", err)
		return nil
	}

	if fi.IsDir() {
		fmt.Println("Directory: ", fn )
	} else {
		fmt.Println("File: ", fn)
	}
	return nil
}