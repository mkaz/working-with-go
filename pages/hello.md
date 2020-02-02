---
title: Hola Mundo
layout: page
order: 1
---

# Hola Mundo in Go

Here is your first go program.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo!")
}
```

All programs must be part of a package, you use `main` as the package for executable programs.

You must import all packages used, including the standard library packages. In this example, the standard library [`fmt`](https://golang.org/pkg/fmt/) is used for formatted I/O.

The standard library is your best friend in Go, it is extensive, often you will not need any third-party dependencies and can rely on just features in the standard library. You can browse all available standard library packages at: <a href="http://golang.org/pkg/">http://golang.org/pkg/</a>.

The main package requires a `main()` function, this is called when the program is run.

Go will run `init()` function before `main()`, if it exists.

Here is an expanded Hello World program that uses a global variable `phrase` defined as a string. <span class="sidenote">All variables have types.</span> The go compiler will automatically detect the type if possible.

```go
package main

import "fmt"

var phrase string

func init() {
	phrase = "Hola Mundo!"
}

func main() {
	fmt.Println(phrase)
}
```

To run the program in your terminal, save to a file `hello.go` and run using `go run`:

```shell
$ go run hello.go
Hola Mundo!
```
