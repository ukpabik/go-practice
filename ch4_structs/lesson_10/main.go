package main

type Membership struct {
	Type             string
	MessageCharLimit int
}
type User struct {
	Membership
	Name string
}

func newUser(name, membershipType string) User {
	membership := Membership{
		MessageCharLimit: 100,
		Type:             membershipType,
	}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	}

	return User{
		Name:       name,
		Membership: membership,
	}
}
