package main

import (
	"fmt"
	"sync"
)

// A send on an unbuffered channel can proceed if a receiver is ready.
// A send on a buffered channel can proceed if there is room in the buffer.
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
