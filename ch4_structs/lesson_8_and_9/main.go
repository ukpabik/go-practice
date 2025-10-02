package main

/**
Empty Structs in Go.
*/

func unused() {
	// Empty anonymous struct
	_ = struct{}{}

	// Named empty struct type
	type emptyStruct struct{}
	_ = emptyStruct{}
}

func main() {
	unused()
}
