package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// context is a container to hold values
	// context can hold timer values or normal values as well
	// we need a container to store context values

	// if no context is available, we will create a new container to store the context values

	// create a new empty context // which is an empty container

	ctx := context.Background()

	// TODO is a context with no values
	// but it gives a signal that we are not sure what context to use , empty or from somewhere else
	//ctx := context.TODO()

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel() // cancel cleanup the resources taken by the context

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://google.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytesData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytesData))

}
