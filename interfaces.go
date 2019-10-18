package main

import "fmt"

type User interface {
	Permissions() int
	Name() string
}

type Admin struct {
	name string
}

func (a Admin) Permissions() int {
	return 5
}

func (a Admin) Name() string {
	return a.name
}

// lets validate if the user is admin. Only admin can login.
func auth(user User) string {
	if user.Permissions() > 4 {
		return user.Name() + " is admin"
	}
	return ""
}

func main() {
	admin := Admin{"Sebas"}
	fmt.Println(auth(admin))
}
