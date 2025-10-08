package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := new(sync.WaitGroup)
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			workV3(i)
		})
	}
	wg.Wait()

}

func workV3(id int) {
	fmt.Println("work done", id)
}
