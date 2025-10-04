package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	userMap := make(map[string]user)
	for i := range len(phoneNumbers) {
		number := phoneNumbers[i]
		name := names[i]
		userMap[name] = user{
			name:        name,
			phoneNumber: number,
		}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
