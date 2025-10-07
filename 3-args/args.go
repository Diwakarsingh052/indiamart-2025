package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	// skip the first argument [ program name ]
	args := os.Args[1:]
	fmt.Println(args)
	//err default value is nil
	_, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println(err)
	}
}
