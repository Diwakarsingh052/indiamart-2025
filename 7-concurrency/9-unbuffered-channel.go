package main

import (
	"fmt"
	"sync"
	"time"
)

// A send on an unbuffered channel can proceed if a receiver is ready.
// A send on a buffered channel can proceed if there is room in the buffer.
func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Go(func() {
		time.Sleep(1 * time.Second)
		fmt.Println(<-ch) // receive
	})

	wg.Go(func() {
		ch <- 10 // send
	})

	wg.Wait()

}
