package main

import (
	"fmt"
	"log"
	"net/http"
)

// http.HandlerFunc
func handler(w http.ResponseWriter, r *http.Request) {
	/**
	An http.Request is a data structure that represents the client HTTP request.
	r.URL.Path is the path component of the request URL.
	The trailing [1:] means "create a sub-slice of Path from the 1st character to the end."
	This drops the leading "/" from the path name.
	**/
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func main() {
	// handle all requests to the root with handler
	http.HandleFunc("/", handler)
	// ListenAndServe specifying it should listen on port 8080 on any interface
	// ListenAndServe always returns an error, since it only returns when something
	// unexpected occurs
	// log it with log.Fatal
	http.HandleFunc("/header", headers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
