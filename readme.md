
# Working with Go

Marcus Kazmierczak [mkaz.com](http://mkaz.com/), [@mkaz](https://twitter.com/mkaz)

Working with Go is a set of example programs in Go (golang) to get an experienced programmer familiar with Go. Note, Go is often referred to as golang to help searches.

This initial started out as a full fledged book, but I found the real value was in all the examples and the text and descriptions were simple enough to include in the code as comments. Also descriptions worked better in full context.


## Install Go

Go is distributed as a binary on the major platforms, Linux, FreeBSD, Mac OS X and Windows. It is available for both 32-bit and 64-bit architectures. See
the official [Getting Started with Go](http://golang.org/doc/install) page for downloads and more instructions.

Go is available in most Linux package environments as `golang`. I would recommend using at least Go 1.1 as the minimum version. So for stable Debian Wheezy 7.x users, 1.0 is latest in stable, you may want to switch to Jessie/Unstable. Most Ubuntu versions have Go 1.2 in repository.


## Clone and Go

Once you have Go installed, clone this repository 

    $ git clone https://github.com/mkaz/working-with-go

And then you can start working through the examples, in any order you want, but they are numbered and build upon themselves. I explain more in the first few and then assume you understand certain packages and structure.

To run the first program use:

    $ go run 01-hello.go

If you want to build a program in Go and then run:

    $ go build -o hello 01-hello.go
    $ ./hello


## Contribute

Working with Go was started by Marcus Kazmierczak, but I welcome and encourage any additional examples, corrections or any contributions to create a good working set of code someone new can use to learn Go.

## Resources

This set of examples assumes a certain level of programming experience and is intended for someone learning the Go language and not someone new to programming altogether.

If you are starting out and want to learn how to program and choose Go as your first language, check out [Learn Programming in Go](http://www.golang-book.com/)


## License

Working with Go is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>. 

