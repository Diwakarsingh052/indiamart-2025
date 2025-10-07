package main

import "fmt"

// Failed Update Mistake Number 1, Nil pointer
var x = 10

// in Go everything is pass by value
// in case of pointers we copy the address from one partner to another
func main() {
	var p *int // nil // default value of a pointer is nil
	fmt.Println(&p, "address of p in the main func")
	updateValue(&p) // copying nil to p1 pointer

	// this will panic
	fmt.Println(*p)
}

func updateValue(p1 **int) {
	fmt.Println(p1, "value of p1 in the updateValue func")
	*p1 = &x   // updated nil pointer to a valid address
	**p1 = 200 // dereference and update the value

}
