package main

import "indiamart-2025/4-slices/inspect"

func main() {
	a := make([]int, 0, 100) // len, cap

	a = append(a, 10)

	inspect.Slice("a", a)

	copy()
}
