package main

import (
	"context"
	"database/sql"
	"strconv"
	"time"
)

// Context must be passed as the first argument
func Slow(ctx context.Context, msg string) (int, error) {
	i, err := strconv.Atoi(msg)
	time.Sleep(2 * time.Second)
	if err != nil {
		return 0, err
	}
	sql.DB{}.ExecContext()
	return i, nil

}
