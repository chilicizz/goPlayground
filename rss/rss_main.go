package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
)

// Channel struct to represent the data
type Channel struct {
	Title           string `xml:"channel>title"`
	Items           []Item `xml:"channel>item"`
	URL             string `xml:"channel>link"`
	PublicationDate string `xml:"channel>pubDate"`
}

// Item in Feed
type Item struct {
	Title           string `xml:"title"`
	URL             string `xml:"link"`
	Author          string `xml:"author"`
	GUID            string `xml:"guid"`
	Description     string `xml:"description"`
	PublicationDate string `xml:"pubDate"`
}

func main() {
	filename := "rss.xml"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Fetching from url")
		url := "https://www.info.gov.hk/gia/rss/general_en.xml"
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		ioutil.WriteFile(filename, data, 0644)
	}
	log.Println("Decoding data")
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	var channel Channel
	err = decoder.Decode(&channel)
	if err != nil {
		log.Fatal("Error Decode: " + err.Error())
	}
	log.Println(channel.Title)

	for _, item := range channel.Items {
		fmt.Println(item.PublicationDate + " " + item.Title)
		//log.Println(item.Description)
	}

	//log.Println("Marshal to JSON")
	//json, _ := json.Marshal(rss)
	//fmt.Printf("%#v\n", json)
}

// parse an item
func ParseItem(data []byte) (Item, error) {
	var item Item
	err := xml.Unmarshal(data, &item)
	return item, err
}

// Parse the channel xml
func Parse(data []byte) (Channel, error) {
	var channel Channel
	err := xml.Unmarshal(data, &channel)
	return channel, err
}
