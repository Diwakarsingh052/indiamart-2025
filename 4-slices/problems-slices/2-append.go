package main

import "indiamart-2025/4-slices/inspect"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}
	inspect.Slice("x", x)

	a := x[2:6]
	inspect.Slice("a", a)

	a = append(a, 888)

	inspect.Slice("a", a)
	inspect.Slice("x", x)

	a = append(a, 999)

	inspect.Slice("a", a)
	inspect.Slice("x", x)

}
