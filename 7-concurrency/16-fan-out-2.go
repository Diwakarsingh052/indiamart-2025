package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wgTask := new(sync.WaitGroup)

	ch := make(chan int, 5)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 6; i++ {
			// keeping track of workers with a dedicated waitgroup
			wgTask.Add(1)
			go func() {
				defer wgTask.Done()
				// fan out pattern
				// spinning n number of goroutines for n tasks
				ch <- i
			}()

		}
		fmt.Println("all values sent")

		wg.Add(1)
		go func() {
			defer wg.Done()
			wgTask.Wait() // waiting for workers to finish
			close(ch)
		}()
		// sends a signal to stop the range

		// once all the tasks are done then we would close the channel
		// close signal range that no more values be sent
		//and it can stop after receiving remaining values

	}()

	wg.Add(1)

	// worker pool pattern,
	// multiple goroutines are working on the same task to receive and process
	go func() {
		defer wg.Done()
		// it would run infinitely, channel needs to be closed to stop this range
		// if channel is closed range will receive remaining values and stop
		for i := range ch { // range over channel is a receive operation
			//time.Sleep(1 * time.Second)
			fmt.Println(i, "range 1")
		}
	}()
	go func() {
		defer wg.Done()
		// it would run infinitely, channel needs to be closed to stop this range
		// if channel is closed range will receive remaining values and stop
		for i := range ch { // range over channel is a receive operation
			//time.Sleep(1 * time.Second)
			fmt.Println(i, "range 1")
		}
	}()

	wg.Wait()
}
