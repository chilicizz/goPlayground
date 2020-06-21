package main

import(
	"fmt"
)


func main() {
	var a [4]int
	a[0] = 1
	var i = a[0]
	fmt.Printf("%s\n", i)
	// i == 1
	// GO's arrays are values, an array variable denotes the entire array.
	// It is not a pointer to the first array element (as in C)
	// this means that when you assign or pass an array value you make a copy
	// of its contents.
	// To avoid this you can pass a pointer to the array

	var b = [2]string{"Penn", "Teller"}
	b = [...]string{"Penn", "Teller"} // compiler can count the number
	fmt.Printf("%s\n", b)

	// Slices
	// slices have no specified length
	var letters = []string{"a", "b", "c", "d"}
	fmt.Printf("%s\n", letters)
	// a slice can be created with the built in function called make
	// func make([]T, len, cap) []T
	var s []byte
	s = make([]byte, 5, 5)
	fmt.Printf("%s\n", s)
	// s == []byte{0,0,0,0,0}
}
