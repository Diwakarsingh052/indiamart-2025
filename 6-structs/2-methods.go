package main

import "fmt"

type Student struct {
	name string
	age  int
}

// method signature
// func (receiver) funcSignature {//body}

func (s Student) Print() {
	fmt.Println(s.name)
}

func main() {
	var s Student
	s.name = "John"
	s.Print()
}
