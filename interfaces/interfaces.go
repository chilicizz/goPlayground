package main

import (
	"fmt"
)

type Whistle string
func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string
func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

type Robot string
func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Walk walk")
}

func (r Robot) String() string {
	// Implement the Stringer interface for printing
	return string(r) + " Robot"
}

func AcceptAnything(a interface{}) {
	// this is an empty interface which means
	// the function can accept any type
	fmt.Println(a)
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()

	toy = Horn("Toyco Blaster")
	toy.MakeSound()

	var toy2 NoiseMaker
	toy2 = Robot("Botco Ambler")
	play(toy2)

	robot, ok := toy2.(Robot) // assert the type
	if ok {
		robot.Walk()
	}
	fmt.Println(robot)
	fmt.Println(toy2)

	AcceptAnything(toy)
}


