package main

import "fmt"

type userV2 struct {
	name string
	age  int
}

func (u userV2) print() {
	fmt.Println(u.name, u.age)
}
func main() {
	//var u1 []user
	u := []userV2{
		{
			name: "abc",
		},
		{
			name: "abc",
			age:  20,
		},
		{
			name: "abc",
			age:  20,
		},
	}

	for _, v := range u {
		v.print()
	}

}
