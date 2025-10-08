package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	go func() {

		// we can't return values from goroutine directly to other goroutine
		id := userId()
		fmt.Println(id)
	}()

	fmt.Println("main finished ")

	wg.Wait()

	//https://go.dev/ref/spec#Send_statements

}

func userId() int {
	return 101
}
