package main

import (
	"fmt"
	"sync"
)

func main() {
	// new returns a pointer to the type, type would be initialized to its default values
	wg := new(sync.WaitGroup)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// anonymous goroutine function, function without name
		go func(id int) {
			defer wg.Done()
			workV2(id)
		}(10) // () this means we are calling the function
	}

	fmt.Println("end of the main")
	wg.Wait()

}

func workV2(id int) {
	fmt.Println("work done", id)
}

func work() {

}
