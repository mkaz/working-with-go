
# Working with Go

By Marcus Kazmierczak [mkaz.com](http://mkaz.com/), [@mkaz](https://twitter.com/mkaz)

Working with Go is a set of example programs in Go (golang) to get an experienced programmer familiar with Go. The Go language is often referred to as golang to help searches.

This initially started out as a full fledged book, but I found the real value was in all the examples; the text and descriptions were simple enough to include in the code as comments. Also I think the descriptions explain better in context of the code.


## Install Go

Go is distributed as a binary on the major platforms, Linux, FreeBSD, Mac OS X and Windows. It is available for both 32-bit and 64-bit architectures. See the official [Getting Started with Go](http://golang.org/doc/install) page for downloads and more instructions.

Go is available in most Linux package environments as `golang`. I would recommend using at least Go 1.1 as the minimum version. So for Debian users, Go 1.0 is latest version in stable, so you may want to switch to Jessie/Unstable or install from source or alternate binary. Most Ubuntu versions have Go 1.2 in repository.


## Clone and Go

Once you have Go installed, clone this repository 

    $ git clone https://github.com/mkaz/working-with-go

And then you can start working through the examples, in any order you want, but they are numbered and build upon themselves. I explain more in the first few and then assume you understand certain packages and structure.

To run the first program use:

    $ go run 01-hello.go

If you want to build a program in Go and then run:

    $ go build -o hello 01-hello.go
    $ ./hello


## Development Environment

Go is a light-weight, which makes it flexible for just about any set of developer tools. A full blown IDE, such as Eclipse, is rarely used. Most developers opt for their favorite text editor and the terminal to run. This is how I run on Linux, but I also split my time on a Mac.

On Mac, I use [TextMate 2](http://macromates.com/download) which has a nice Go bundle. The bundle offers a couple of shortcuts, the two I use frequently are  quick running of program <b>&#8984;R</b> and the shortcut to reformat code <b>^&#8679;H</b>. 

If using <b>vim</b>, I recommend using <a
href="https://github.com/fatih/vim-go">vim-go</a> package. 

Go has two tools that can automatically format your source code to the Go
coding standard, `gofmt` which comes with standard install and `goimports`
which will format and auto adjust import statements as needed.

I recommend installing and configuring <a
href="https://godoc.org/golang.org/x/tools/cmd/goimports">goimports</a> to automatically run on save.

Install goimports:

    $ go get golang.org/x/tools/cmd/goimports


Setup vim to auto run `goimports` on save using vim-go:

    let g:go_fmt_command = "goimports"



## Contribute

Working with Go was started by Marcus Kazmierczak. Additions, corrections and any contributions are encouraged, please submit a pull request with your change or an issue for a bug or fix. Thanks to <a href="https://github.com/mkaz/working-with-go/graphs/contributors">all the contributors</a>!

## Resources

This set of examples assumes a certain level of programming experience and is intended for someone learning the Go language and not someone new to programming altogether.

If you are starting out and want to learn how to program and choose Go as your first language, check out [Learn Programming in Go](http://www.golang-book.com/)

The official site has a [Tour of Go](http://tour.golang.org/) which is an interactive walk through, another good introduction to the language.


## License

Working with Go is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>. 

