package main

import (
	"fmt"
	"sync"
	"time"
)

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
