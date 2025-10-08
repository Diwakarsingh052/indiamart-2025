package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1) // add 1 to the counter
	go helloV2(&wg)

	fmt.Println("doing some important work")

	wg.Wait() // block until the counter is zero
	fmt.Println("end of the main")
}

func helloV2(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("hello")
	wg.Done() // done decrements the counter by one
}
