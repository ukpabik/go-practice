package main

/**
Embedded Structs in Go.
*/

type sender struct {
	user
	rateLimit int
}

type user struct {
	name   string
	number int
}
