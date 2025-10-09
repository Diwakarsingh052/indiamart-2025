package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/text", Mid(homeV2))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func homeV2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
	fmt.Println("homeV2")
}

func Mid(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("middleware outside layer")
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware preprocessing")
		next(w, r)
		fmt.Println("middleware postprocessing")

	}
}

//func Mid2(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("middleware 2 preprocessing")
//		next(w, r)
//		fmt.Println("middleware 2 postprocessing")
//
//	}
//}

//https://codeshare.io/G6ddLN
