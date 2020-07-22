package main

import (
	"fmt"
)
type subscriber struct {
	name	string
	rate	float64
	active	bool
}

func printInfo(s *subscriber) {
	// Passing in the subscriber as a pointer
	// otherwise this is pass by value
	fmt.Println("Name:", (*s).name)
	// dereference the pointer to get s
	// However struct.field directly is also ok
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?:", s.active)
}

func main() {
	fmt.Println("Structs")

	// defining a struct
	var myStruct struct {
		number  float64
		word    string
		toggle  bool
	}
	fmt.Printf("%#v\n", myStruct)
	myStruct.number = 3.14
	fmt.Println(myStruct.number)

	// type definition
	/**
	type myType struct {
		// fields 
	}
	**/

	var subscriber1 subscriber
	subscriber1.name = "Aman Singh"
	fmt.Println("Name:", subscriber1.name)

	// define as a literal
	var subscriber2 subscriber = subscriber{name : "Joy Carr", rate : 4.5, active : true}
	// functions pass structs by value -> ie copy
	// passing in the address as a reference
	printInfo(&subscriber1)
	printInfo(&subscriber2)
	// for packages you would need to capitalise the first letter to export the struct
	// this is the same for the field

}
