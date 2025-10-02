package main

import "fmt"

/**
Implementing multiple interfaces in Go.
*/

func (e email) cost() int {
	// ?
	var cost int
	if e.isSubscribed {
		cost = 2 * len(e.body)
	} else {
		cost = 5 * len(e.body)
	}
	return cost
}

func (e email) format() string {
	// ?
	var subscribeMessage string
	if e.isSubscribed {
		subscribeMessage = "Subscribed"
	} else {
		subscribeMessage = "Not Subscribed"
	}

	return fmt.Sprintf("'%s' | %s", e.body, subscribeMessage)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
