package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
)

type Post struct {
	Id	int
	Content	string
	Author	string
}

func main () {
	// create a new file
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello world!", Author: "Cyril"},
		Post{Id: 2, Content: "Bonjour monde!", Author: "Frenchie"},
	}

	// create the writer
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	// Reading a file
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file) // csv reader
	reader.FieldsPerRecord = -1 // -> we don't care if there are missing fields
	// if positive number then we expect that number of fields otherwise error
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// recreate the object by parsing the output
	var posts[]Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
