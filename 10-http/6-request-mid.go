package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// The provided key must be comparable and should not be of type string or
// any other built-in type to avoid collisions between packages using context.
// Users of WithValue should define their own types for keys.
type ctxKey string

const reqIdKey ctxKey = "reqId"

func main() {

	http.HandleFunc("/hello", ReqIdMid(Hello))
	http.ListenAndServe(":8080", nil)

}

func ReqIdMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// adding the reqId to the context
		// context stores the key value pair
		ctx := r.Context()
		ctx = context.WithValue(ctx, reqIdKey, uuid.NewString())

		// updating the request object with the new context
		next(w, r.WithContext(ctx))
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// type assertion
	// check what is the type of the value stored in the interface
	// use ok variant to avoid panic
	reqId, ok := ctx.Value(reqIdKey).(string)
	if !ok {
		reqId = "unknown"
	}
	fmt.Println(reqId, "hello")
	InternalProcess(reqId)
	w.Write([]byte("hello"))
}

func InternalProcess(reqId string) {
	fmt.Println("internal process", reqId)
	// access uuid here
}
