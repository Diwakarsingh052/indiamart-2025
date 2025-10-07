package main

import (
	"fmt"
)

type FileDetails struct {
	name string
}

func (f *FileDetails) PrintName() {
	fmt.Println(f.name)
}
func (f *FileDetails) UpdateName(fileName string) {
	// nil pointer exception would happen if f is nil
	f.name = fileName
}

func main() {
	var f1 *FileDetails
	// f1 is nil pointer
	f1.UpdateName("abc")
	f1.PrintName()

}
