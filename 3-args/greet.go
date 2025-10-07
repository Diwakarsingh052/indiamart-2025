package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	greet()
}

func greet() {
	if len(os.Args) != 4 {
		fmt.Println("wrong number of arguments")
		return
	}

	userParams := os.Args[1:]

	name := userParams[0]
	age, err := strconv.Atoi(userParams[1])
	if err != nil {
		fmt.Println("Error converting age to int. Error message: ", err)
		return
	}
	marks, err := strconv.Atoi(userParams[2])
	if err != nil {
		fmt.Println("Error converting marks to int. Error message: ", err)
		return
	}

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("marks: ", marks)
}
