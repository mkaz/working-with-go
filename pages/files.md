---
title: Files
date: 2014-06-09
order: 9
---

# Working with Files


## Reading File

You can read a complete file in using [`ioutil.ReadFile`](https://golang.org/pkg/io/ioutil/#ReadFile). The function returns two variables, the first the content of the file. The second variable is the error if one occurred. If no error, than `err` will be `nil`

```go
filename := "./extras/rabbits.txt"

content, err := ioutil.ReadFile(filename)
if err != nil {
	log.Fatalln("Error reading file", filename)
}
```

The content is returned as a `[]byte` and not a string. You need to cast to a string to use as such, for example to display. Use `string()` to cast a `[]byte` to `string` type.

```go
fmt.Println(string(content))
```

## Check if File Exists

One of the above errors can be if the file does not exist. You can use [`os.Stat`](https://golang.org/pkg/os/#Stat) to check explicitly if file exists without trying to open.

```go
if _, err := os.Stat("junk.foo"); os.IsNotExist(err) {
	fmt.Println(">>>")
	fmt.Println("File: junk.foo does not exist")
}
```

## Write to a new File

Use [`ioutil.WriteFile`](https://golang.org/pkg/io/ioutil/#WriteFile) to write a file out. The function takes three variables, the filename, the content (as a `[]byte`) and the file system mode.

```go
outfile := "output.txt"
err = ioutil.WriteFile(outfile, content, 0644)
if err != nil {
	fmt.Println("Error writing file: ", err)
} else {
	fmt.Println(">>>")
	fmt.Println("Created: ", outfile)
}
```

## Append to an existing File

You can write out to an existing file appending the content using the following.

```go
af, err := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY, 0644)
if err != nil {
	fmt.Println("Error appending to file:", err)
}
defer af.Close()
if _, err = af.WriteString("Appending this text"); err != nil {
	fmt.Println("Error writing to file:", err)
}
```

The `defer` statement defers the execution until the surrounding function (or overall program) completes. You should always use `defer` for something that needs to be closed, or cleaned up.

## Using Filepath

Use the [filepath package](https://golang.org/pkg/path/filepath/) for working with cross-platform paths properly.
For example, use `filepath.Join` for creating a path with directory.

```go
package main

import (
    "fmt"
    "path/filepath"
)

func main() {
    fmt.Println(filepath.Join("a", "b", "file.ext"))
}
```

See filepath package documentation for additional function, including splitting paths, checking filename extension, base and more.
