package main

import (
	"fmt"
	"sync"
)

func main() {
	get := make(chan string, 1)
	post := make(chan string, 1)
	put := make(chan string, 1)

	wg := new(sync.WaitGroup)
	wg.Go(func() {
		//time.Sleep(3 * time.Second)
		get <- "get"

	})

	wg.Go(func() {
		//time.Sleep(1 * time.Second)
		post <- "post"
	})

	wg.Go(func() {
		put <- "put"
	})

	//fmt.Println(<-get)
	//fmt.Println(<-post)
	//fmt.Println(<-put)
	for i := 1; i <= 3; i++ {
		//whichever case is not blocking exec that first
		//whichever case is ready first, exec that.
		// possible cases are chan recv , send , default

		select {
		case x := <-get:
			fmt.Println(x)
		case x := <-post:
			fmt.Println(x)
		case x := <-put:
			fmt.Println(x)
		}
	}
	wg.Wait()
}
