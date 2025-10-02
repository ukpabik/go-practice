package main

import "fmt"

/**
Struct Methods in Go.
*/

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}
