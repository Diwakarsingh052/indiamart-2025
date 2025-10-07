package main

import "fmt"

func main() {
	// slice is a three word data structure
	// pointer to an array
	// length
	// capacity

	x := []int{}
	var a []int //nil
	//a[0] = 10 // updating the value at index 0 // not appending
	i := []int{10, 20, 30, 40}

	//runtime error: index out of range
	//would not be detected at compile time
	//i[100] = 1000

	fmt.Println(i)
	fmt.Println(a)

	//%#v give more info about the data
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", x)

	for i := 0; i < 10; i++ {

	}

	for index, value := range i {
		fmt.Println(index, value)
	}

}
