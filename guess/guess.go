package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	maxGuesses := 10
	target := rand.Intn(100) + 1

	reader := bufio.NewReader(os.Stdin)
	var success bool = false

	fmt.Print("Guess a number between 1 and 100: ")
	for i := 0; i < maxGuesses; i++ {

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Your guess is too low")
		} else if guess > target {
			fmt.Println("Your guess is too high")
		} else {
			fmt.Println("Congratulations, you guessed correctly!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Sorry you failed to guess, the correct number is:", target)
	}
}
