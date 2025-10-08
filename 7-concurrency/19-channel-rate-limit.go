package main

import (
	"fmt"
	"sync"
	"time"
)

// MaxGoroutines is the maximum number of concurrent goroutines
const MaxGoroutines = 2

func main() {
	wg := new(sync.WaitGroup)
	sem := make(chan struct{}, MaxGoroutines)

	for i := 1; i <= 10; i++ {
		sem <- struct{}{}
		wg.Go(func() {
			defer func() { <-sem }()
			fmt.Println("working on work id", i)
			time.Sleep(1 * time.Second)
		})
	}

	wg.Wait()

}
