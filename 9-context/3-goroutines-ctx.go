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
		i := slowFunc(ctx)
		select {
		case <-ctx.Done():
			fmt.Println("send cancelled,", ctx.Err())
			return
		case ch <- i:
			fmt.Println("sent", i)
		}

	})

	// select would block until either of the cases is ready
	select {
	// if timeout happens, then the receiver would move on
	case <-ctx.Done():
		fmt.Println(ctx.Err())

		// receiving the value normally
	case x := <-ch:
		fmt.Println("received", x)

	}

	wg.Wait()

}

func slowFunc(ctx context.Context) int {
	// use select to check if the context is cancelled
	// if cancelled, then rollback the changes
	// use simple print statements to show the flow of the program

	time.Sleep(2 * time.Second)
	fmt.Println("slowFunc() ran and changed 10 files")
	select {
	case <-ctx.Done():
		fmt.Println("rollback the changes")
		return 0
	default:
	}
	return 10

}
