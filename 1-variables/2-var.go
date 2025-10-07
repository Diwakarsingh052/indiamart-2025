package main

import "fmt"

func main() {
	//var firstName string // camelCase for variable naming
	// every types have a default value
	var a int // int default is 0
	var b string = "hello"
	var c = "rahul"
	//c = 1 // type mismatch
	fmt.Println(a, b, c)
	// short hand can't be used at global level
	// go compiler would infer the type automatically from the right side value
	d, ok := 10, true // shorthand// create and assign
	_, _ = d, ok
}
