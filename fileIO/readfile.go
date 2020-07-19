package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	// Returns the pointer to the file and error
	if err != nil {
		log.Fatal(err)
	}
	// pass the file to bufio to return a Scanner
	scanner := bufio.NewScanner(file)

	// Scanner Scan() function should be used in for loop
	// returns true if it read data
	for scanner.Scan() {
		// loop until end of file and print out the text
		fmt.Println(scanner.Text())
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
