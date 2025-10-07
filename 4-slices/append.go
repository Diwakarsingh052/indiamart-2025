package main

import (
	"fmt"
	"indiamart-2025/4-slices/inspect"
)

//If the capacity of s is not large enough to fit the additional values,
//append allocates a new, sufficiently large underlying array
//that fits both the existing slice elements and the additional values.
//Otherwise, append re-uses the underlying array.

func main() {

	s := []int{1}
	fmt.Println(len(s), cap(s), &s[0])

	s = append(s, 10, 20)
	inspect.InspectSlice("s", s)
	s = append(s, 100, 200)
	inspect.InspectSlice("s", s)
	s = append(s, 300)
	inspect.InspectSlice("s", s)
}
