package main

import (
	"fmt"
	"sync"
)

var x = 1

func main() {
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	wg.Go(func() {
		UpdateX(10, m)
	})
	wg.Go(func() {
		UpdateX(20, m)
	})
	wg.Go(func() {
		PrintX(m)
	})

	wg.Wait()

}
func PrintX(m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	fmt.Println(x)
}
func UpdateX(val int, m *sync.Mutex) {
	// shared resources
	/*
	   1. Global Variables
	   2. Pointer variables
	   3. Structs fields where methods are pointers
	   4. Maps,
	*/

	// data race situations
	//	- at least one concurrent write and n number of reads
	//	- n number of concurrent writes
	// 	- n number of concurrent writes and n number of concurrent reads
	// 	Note - Data race doesn't happen if there are only concurrent reads

	// run program with race detector
	// go run -race main.go // don't use this in production
	m.Lock()
	defer m.Unlock()
	x = val

}
