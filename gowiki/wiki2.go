package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

// Data Structure
type Page struct {
	Title string
	Body []byte // Use byte[] a byte slice as it is expected by the io libraries
}

// This is a method named save that takes as its receiver p, a pointer to Page. 
// It takes no parameters and returns a value of type error
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// The function loadPage constructs the file anme from the title parameter,
// reads the file's contents into a new variable body, and returns a pointer to
// a Page literal constructed with the proper title and body values
// the function ReadFile can return multiple values, 
// _ is the blank identifier which throws away the error
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	// body, _ := ioutil.ReadFile(filename)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>" + 
	"<form action=\"/save/%s\" method=\"POST\">" + 
	"<textarea name=\"body\">%s</textarea><br>" + 
	"</form>", p.Title, p.Title, p.Body)
}

func main() {
	/**
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
	*/
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
