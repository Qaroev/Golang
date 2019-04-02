package main

type User struct {
	name  string
	admin bool
}

func (u User) IsAdmin() bool {
	return u.admin
}
