package main

import (
	"fmt"
	"sync"
)

func main() {
	// new returns a pointer to the type, type would be initialized to its default values
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("work done")
	}()

	fmt.Println("end of the main")
	wg.Wait()

}
