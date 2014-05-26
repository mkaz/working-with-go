/**
 * web-server.go
 *
 * A program example web server
 *
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// define function to handle requests
	http.HandleFunc("/hello", helloRequest)
	http.HandleFunc("/", getRequest)

	// start web server
	fmt.Println("Listening on http://localhost:9999/")
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// basic handler for /hello request
func helloRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hallo Welt")
	return
}

// setup as default file server
func getRequest(w http.ResponseWriter, r *http.Request) {
	file_requested := "./" + r.URL.Path
	http.ServeFile(w, r, file_requested)
	return
}
