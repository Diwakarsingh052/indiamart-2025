package main

import (
	"fmt"
	"sync"
)

// BAD CODE AHEAD
// don't use infinite for select with buffered channels
// we can lose data
// if you want to use a buffered channel,
// then run goroutines and run range inside them to receive values
func main() {
	get := make(chan string, 3)
	post := make(chan string, 3)
	put := make(chan string, 3)

	done := make(chan struct{})

	wg := new(sync.WaitGroup)
	wgTask := new(sync.WaitGroup)

	// buffered channel don't care about receivers
	// in this code if done channel is closed before receiving
	//then we would lose values

	wgTask.Go(func() {
		//time.Sleep(3 * time.Second)
		get <- "get"
		get <- "get2"
		get <- "get3"

	})

	wgTask.Go(func() {
		//time.Sleep(1 * time.Second)
		post <- "post"
	})

	wgTask.Go(func() {
		put <- "put"
	})

	wg.Go(func() {
		wgTask.Wait()
		close(done)
	})

	wg.Go(func() {

		for {
			select {
			case x := <-get:
				fmt.Println(x)
			case x := <-post:
				fmt.Println(x)
			case x := <-put:
				fmt.Println(x)
			case <-done:
				fmt.Println("all goroutines done")
				return
			}
		}
	})

	//fmt.Println(<-get)
	//fmt.Println(<-post)
	//fmt.Println(<-put)

	wg.Wait()
}
