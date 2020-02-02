---
title: Directories
layout: page
order: 10
---

# Directories

## Read Directory Recursively

Use the [`filepath.Walk`](https://golang.org/pkg/path/filepath/#Walk) function to recursively walk down a directory.

The `Walk` function accepts a path, and [`WalkFunc`](https://golang.org/pkg/path/filepath/#WalkFunc) as its parameters.

The `WalkFunc` signature is `func(path string, info os.FileInfo, err error) error`, below is an example using an anonymous function with that signature.

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

	filepath.Walk(path, func(fn string, fi os.FileInfo, err error) error {
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
    })
}
```

## Create Directory if not exists

Check if directory exists, create directory if it does not.

```go
    pathdir := filepath.Join(path, dir)
    if _, err := os.Stat(pathdir); os.IsNotExist(err) {
        mdErr := os.Mkdir(dirpath, 0755)
        if mdErr != nil {
            fmt.Println("Error making directory", mdErr)
        }
    }
```

## Get User Home Directory

A platform independent way to get the user's home directory.

```go
import "os/user"

usr, err := user.Current()
fmt.Println("User Home Directory:", usr.HomeDir)
```

