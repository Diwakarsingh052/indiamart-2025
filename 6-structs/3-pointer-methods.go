package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u1 *user) UpdateName(name string) {
	// dereference of a pointer happens automatically in structs
	u1.name = name
}

func main() {
	u := user{name: "John", age: 20}
	u.UpdateName("Jane")
	fmt.Println(u)
}
