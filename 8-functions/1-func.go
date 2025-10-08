package main

import "fmt"

func main() {
	calc(add)
	calc(sub)
	calc(func(a, b int) {
		fmt.Println(a * b)
	})
}

func add(a, b int) {
	fmt.Println(a + b)
}
func sub(a, b int) {
	fmt.Println(a - b)
}

// calc function would accept any function as an argument
// whose signature is matches with func(int, int)
func calc(f func(int, int)) {
	f(1, 2)
}
