package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func displayHeader(writer http.ResponseWriter, request *http.Request) {
	h := request.Header
	fmt.Fprintln(writer, h)
}

func displayBody(w http.ResponseWriter, r *http.Request) {
	size := r.ContentLength
	body := make([]byte, size)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func notYetImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "Not yet implemented")
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

type Post struct {
	User string
	Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "Example User",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/displayHeader", displayHeader)
	http.HandleFunc("/displayBody", displayBody)
	http.HandleFunc("/unimplemented", notYetImplemented)
	http.HandleFunc("/redirect", redirectHandler)
	http.HandleFunc("/json", jsonExample)
	http.ListenAndServe(":8080", nil)
}
