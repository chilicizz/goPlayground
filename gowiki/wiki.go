package main

import (
	"fmt"
	"io/ioutil"
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


func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
