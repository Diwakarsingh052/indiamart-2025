package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	i, err := Slow(ctx, "10")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("converted", i)

}

// Context must be passed as the first argument
func Slow(ctx context.Context, msg string) (int, error) {
	i, err := strconv.Atoi(msg)
	time.Sleep(2 * time.Second)
	if err != nil {
		return 0, err
	}
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		// rollback your previous work if you don't want to persist the changes
	}
		fmt.Println("client is still alive")
	}

	return i, nil

}
