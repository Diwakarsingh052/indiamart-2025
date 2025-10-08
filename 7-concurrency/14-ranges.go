package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)

	ch := make(chan int, 5)
	wg.Go(func() {
		for i := 1; i <= 6; i++ {
			ch <- i
		}
		fmt.Println("all values sent")
		// sends a signal to stop the range
		close(ch)
		// close signal range that no more values be sent
		//and it can stop after receiving remaining values
		//ch <- 10 // once the channel is closed, we can't send more values to it

	})

	wg.Go(func() {

		// it would run infinitely, channel needs to be closed to stop this range
		// if channel is closed range will receive remaining values and stop
		for i := range ch { // range over channel is a receive operation
			time.Sleep(1 * time.Second)
			fmt.Println(i, "range 1")
		}
	})

	wg.Wait()
}
