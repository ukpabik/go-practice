package main

import "unicode"

func isValidPassword(password string) bool {

	if len(password) < 5 || len(password) > 12 {
		return false
	}
	upper := false
	digit := false
	for _, ch := range password {
		if unicode.IsDigit(ch) {
			upper = true
		}
		if unicode.IsUpper(ch) {
			digit = true
		}
	}

	if digit && upper {
		return true
	}
	return false
}
