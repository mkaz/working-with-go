---
title: Command-line
layout: page
order:13
---

# Command-line Args

## os.Args

Command-line args are stored as a slice in `os.Args`. The first argument in list is program itself.

```go
num_args := len(os.Args)
```

You can check to see if any command-line arguments were passed in

```go
if num_args < 2 {
	fmt.Println(">> No args passed in")
}
```

## Flag Library

Use the [`flag`](https://golang.org/pkg/flag/) standard library to parse command-line arguments. The flag library parses parameters passed in and flags for invalid types, invalud flags, and handles most needs for creating a command-line program.

```go
var s string
flag.StringVar(&s, "str", "default value", "text description")
```

Usage:

```shell
$ ./program -str "Hello"
```

## Full Example

Here is a full example program showing how to use command-line arguments and flags in Go

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

// global vars
var str string
var num int
var help bool

func main() {

	// define flags
	flag.StringVar(&str, "str", "default value", "text description")
	flag.IntVar(&num, "num", 5, "text description")
	flag.BoolVar(&help, "help", false, "Display Help")

	// parse
	flag.Parse()

	// check if help was called explicitly
	if help {
		fmt.Println(">> Display help screen")
		os.Exit(1)
	}

	// See values assigned
	fmt.Println(">> String:", str)
	fmt.Println(">> Number:", num)

	// last command-line argument
	fmt.Println(">> Last Item:", os.Args[num_args-1])

	// the os.Args will include flags for example
	// go run command-line-args.go --str=Foo filename
	// os.Args[1] = "--str=Foo"

	// If you have flags and want just the arguments
	// then after calling flag.Parse() you can call
	// flag.Args which store only the args
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println(">> Flag Arg:", args[0])
	}

}
```

Run with different parameters

```shell
$ go run command-line.go
>> No args passed in
>> String: default value
>> Number: 5
>> Last Item: command-line.go

$ go run command-line.go --str=Foo --num=8 filename
>> String: Foo
>> Number: 8
>> Last Item: filename
>> Flag Arg: filename
```
