package main

import (
	"fmt"
)

type user struct {
	name string
	email string
}

func (u *user) notify(){
	fmt.Printf("user name %s, email %s \n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func main(){
	lili := &admin{
			user: user{
				name: "lili",
				email: "lili@lili.com",
			},
			level: "one",
		}
	lili.user.notify()

	lili.notify()
}
