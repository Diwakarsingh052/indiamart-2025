package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// sync.WaitGroup is usable in defualt state
	//wg := &sync.WaitGroup{}

	wg := new(sync.WaitGroup) // exactly the same as above

	wg.Add(1) // add 1 to the counter
	go helloV2(wg)

	fmt.Println("doing some important work")

	wg.Wait() // block until the counter is zero
	fmt.Println("end of the main")
}

func helloV2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("hello")

	// done decrements the counter by one
}
