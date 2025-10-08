package main

import (
	"fmt"
	"sync"
)

func main() {
	get := make(chan string)
	post := make(chan string)
	put := make(chan string)
	done := make(chan struct{})

	wg := new(sync.WaitGroup)
	wgTask := new(sync.WaitGroup)

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
