---
title: Read Directory
layout: post
date: 2014-05-10
---

# Read Directory Recursively

Use the [`filepath.Walk`](https://golang.org/pkg/path/filepath/#Walk) function to recursively walk down a directory.

The `Walk` function accepts a path, and [`WalkFunc`](https://golang.org/pkg/path/filepath/#WalkFunc) as its parameters. The `WalkFunc` signature is `func(path string, info os.FileInfo, err error) error`


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
	} else {
		fmt.Println("File: ", fn)
	}
	return nil
}
```
