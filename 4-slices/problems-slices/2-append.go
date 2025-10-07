package main

import "indiamart-2025/4-slices/inspect"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}
	inspect.InspectSlice("x", x)

	a := x[2:6]
	inspect.InspectSlice("a", a)

	a = append(a, 888)

	inspect.InspectSlice("a", a)
	inspect.InspectSlice("x", x)

	a = append(a, 999)

	inspect.InspectSlice("a", a)
	inspect.InspectSlice("x", x)

}
