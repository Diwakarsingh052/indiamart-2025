package main

import "fmt"

type Person struct {
	name    string // fields of struct
	age     int
	address string
}

func main() {
	var p Person
	p1 := Person{
		name:    "abc",
		age:     30,
		address: "abc",
	}
	_ = p1
	p.name = "rahul"
	fmt.Println(p)
	fmt.Printf("%+v", p)
}
