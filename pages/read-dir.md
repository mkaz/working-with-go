---
title: Read Directory
parent: 1465
template: page-tut.php
order: 10
---

# Read Directory Recursively

Use the [`filepath.Walk`](https://golang.org/pkg/path/filepath/#Walk) function to recursively walk down a directory.

The `Walk` function accepts a path, and [`WalkFunc`](https://golang.org/pkg/path/filepath/#WalkFunc) as its parameters.

The `WalkFunc` signature is `func(path string, info os.FileInfo, err error) error`

In the `WalkFunc` if you want to skip an entire directory, return `filepath.SkipDir` which is a special error variable defined in the library.

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path := "./dir"
	filepath.Walk(path, Walker)
}

func Walker(fn string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("Walker Error: ", err)
		return err
	}

	if fi.IsDir() {
		fmt.Println("Directory: ", fn)
		if fi.Name() == "skipme" {
			return filepath.SkipDir
		}
	} else {
		fmt.Println("File: ", fn)
	}
	return nil
}
```
