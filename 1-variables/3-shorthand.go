package main

import "fmt"

func main() {
	id := 101
	var data int
	if id > 0 {
		// variable shadowing
		// data is a new variable created inside this block
		// below line would not update the outside data variable
		data := loadData() // remove the : to fix the issue
		fmt.Println(data)
	}

	fmt.Println(data)

}

func loadData() int {
	return 100
}
