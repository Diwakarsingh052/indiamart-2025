package main

import (
	"sync"
	"time"
)

func main() {
	get := make(chan string)
	post := make(chan string)
	put := make(chan string)

	wg := new(sync.WaitGroup)
	wg.Go(func() {
		time.Sleep(3 * time.Second)
		get <- "get"
	})

	wg.Go(func() {
		time.Sleep(1 * time.Second)
		post <- "post"
	})

	wg.Go(func() {
		put <- "put"
	})
}
