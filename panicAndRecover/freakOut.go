package main

import ("fmt")

func calmDown() {
	// recover function stops the panic
	// recover returns the panic
	p := recover()
	err, ok := p.(error) // can assert the type of the panic
	if ok {
		fmt.Println("This is an error", err)
	}
	fmt.Println("Unknown panic:", p)
}

func freakOut() {
	// defer the calmDown function to be called 
	// when then function panics
	defer calmDown()
	panic("oh no")
	fmt.Println("This is not run")
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}

