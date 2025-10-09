package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	wg.Go(func() {
		i := slowFunc()
		ch <- i
	})

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case x := <-ch:
		fmt.Println("received", x)

	}

	wg.Wait()

}

func slowFunc() int {
	time.Sleep(2 * time.Second)
	fmt.Println("slowFunc() ran and changed 10 files")
	return 10

}
