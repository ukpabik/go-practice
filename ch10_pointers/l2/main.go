package main

import "strings"

func removeProfanity(message *string) {
	// ?
	*message = strings.ReplaceAll(*message, "fubb", strings.Repeat("*", len("fubb")))
	*message = strings.ReplaceAll(*message, "shiz", strings.Repeat("*", len("shiz")))
	*message = strings.ReplaceAll(*message, "witch", strings.Repeat("*", len("witch")))
}
