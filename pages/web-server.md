---
title: Web Server
layout: page
order: 17
---

# Web Server

The standard library [`net/http `](https://golang.org/pkg/net/http/) provides a powerful set of tools you can use to make your own web server.

See my Lanyon project: [https://github.com/mkaz/lanyon](https://github.com/mkaz/lanyon) for a more complete server created in Go which converts markdown to html.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// associate URLs requested to functions that handle requests
	http.HandleFunc("/hello", helloRequest)
	http.HandleFunc("/", getRequest)

	// start web server
	log.Println("Listening on http://localhost:9999/")
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// basic handler for /hello request
func helloRequest(w http.ResponseWriter, r *http.Request) {

	// Fprint writes to a writer, in this case the http response
	fmt.Fprint(w, "Hallo Welt")
	return
}

// this function simply serves up the file requested, or
// an index list if a directory is requested
func getRequest(w http.ResponseWriter, r *http.Request) {
	file_requested := "./" + r.URL.Path
	http.ServeFile(w, r, file_requested)
	return
}
```
