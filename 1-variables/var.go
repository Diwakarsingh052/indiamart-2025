package main

import "fmt"

// GOOS=linux GOARCH=amd64 go build var.go
// go run var.go // to run programs while developing
func main() {

	var a int = 10
	fmt.Println(a)

}
