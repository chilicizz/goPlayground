package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

type Page struct {
	URL string
	Size int
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)} // return the value to the channel when complete
	//fmt.Println(len(body))
}

func main() {
	pages := make(chan Page)// create the channel
	urls := []string{"https://example.com",
			"https://golang.org",
			"https://golang.org/doc"}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	//go responseSize("https://example.com", sizes)
	//go responseSize("https://golang.org", sizes)
	//go responseSize("https://golang.org/doc", sizes)
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
	//fmt.Println(<-sizes) // print the value when the channel is assigned a value
	//fmt.Println(<-sizes)
	//fmt.Println(<-sizes)
//	go a()
//	go b()
}
