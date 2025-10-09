package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func main() {
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	// if panic happens in http handler function, http will automatically recover the panic
	go func() {
		// if panic happens in goroutine, it will not recover the panic
		// we need to recover the panic manually
		// otherwise service would crash
		defer recoveryFunc()
		var i []int
		fmt.Println(i[0])
	}()
	w.Write([]byte("Hello World"))
}

func recoveryFunc() {
	msg := recover()
	if msg != nil {
		fmt.Println(msg)
		fmt.Println("Recovered in f")
		fmt.Println(string(debug.Stack()))
		return
	}

}
