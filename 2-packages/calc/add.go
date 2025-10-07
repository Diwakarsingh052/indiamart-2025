package calc

// If function first letter is capital, it is exported
// otherwise it is not exported, which means it can only be used in the same package

// Exporting means a type or function can be used outside the package

func Add(a int, b int) {
	sum := a + b

	//unexported function
	// can be used in the same package
	print(sum)
}
