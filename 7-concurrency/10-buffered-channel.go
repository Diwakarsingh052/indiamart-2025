package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 1)
	wg := new(sync.WaitGroup)
	wg.Go(func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("go routine running")
		}
	})
	ch <- 10
	ch <- 20

	wg.Wait()

}
