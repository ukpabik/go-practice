package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	error := fmt.Sprintf("SMS that costs $%.2f to be sent to '%s' can not be sent", cost, recipient)

	return error
}
