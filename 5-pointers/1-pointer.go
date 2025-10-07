package main

import "fmt"

func main() {
	// default value of pointer is nil
	var p *int // pointer to int
	var a int = 10
	p = &a
	// print the address of a
	fmt.Println(&a) // & is the address operator

	// * gives value at the address
	*p = 100 // dereference the pointer
	update(p)
	fmt.Println(a)
}

func update(p *int) {
	*p = 200
}
