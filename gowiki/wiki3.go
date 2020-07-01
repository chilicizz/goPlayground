package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// Data Structure
type Page struct {
	Title string
	Body  []byte // Use byte[] a byte slice as it is expected by the io libraries
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
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/**
If the title is valid, it will be returned along with a nil error value.
If the title is invalid, the function will write a "404 Not Found" error to
the HTTP connection, and return an error to the handler. To create a new error,
we have to import the errors package.
**/
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		// The http.Redirect function adds an HTTP status code of http.StatusFound (302)
		// and a Location header to the HTTP response.
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
	// the .Title and .Body identifiers refer to the p.Title and p.Body
	// Template directives are enclosed in double curly braces.
	// The printf "%s" .Body instruction is a function call that outputs .Body as
	// a string instead of a stream of bytes, the same as a call to fmt.Printf.
	// The html/template package helps guarantee that only safe and correct-looking
	// HTML is generated by template actions. For instance, it automatically escapes
	// any greater than sign (>), replacing it with &gt;, to make sure user data
	// does not corrupt the form HTML.
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	// The page title (provided in the URL) and the form's only field,
	// Body, are stored in a new Page. The save() method is then called
	// to write the data to a file, and the client is redirected to the /view/ page.
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	// The value returned by FormValue is of type string. We must convert that value
	// to []byte before it will fit into the Page struct. We use []byte(body) to
	// perform the conversion.
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// The function template.Must is a convenience wrapper that panics when passed
// a non-nil error value, and otherwise returns the *Template unaltered. A panic
// is appropriate here; if the templates can't be loaded the only sensible thing
// to do is exit the program.
// The ParseFiles function takes any number of string arguments that identify our
// template files, and parses those files into templates that are named after the
// base file name. If we were to add more templates to our program, we would add their
// names to the ParseFiles call's arguments.

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// The function regexp.MustCompile will parse and compile the regular expression,
// and return a regexp.Regexp. MustCompile is distinct from Compile in that it will
// panic if the expression compilation fails, while Compile returns an error as a
// second parameter.

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
