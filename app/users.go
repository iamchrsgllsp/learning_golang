package app

import "fmt"

type Users struct {
	name     string
	password string
	premium  bool
	active   string
}

var allusers = []Users{}

func CreateUser(name string, password string, active string) Users {
	user := Users{name, password, false, active}
	output := append(allusers, user)
	fmt.Println(output)
	return user
}
