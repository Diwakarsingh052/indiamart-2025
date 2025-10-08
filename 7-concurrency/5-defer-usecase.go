package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("abc.txt")
	// always check for error the error first
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	_ = f
}
