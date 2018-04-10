package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name	string
	email	string
}

func (u *user) notify(){
	fmt.Printf("name %s email %s \n", u.name, u.email)
}

type admin struct {
	user
	level	string
}

func main(){
	ad := admin{
		user: user{
			name: "ad",
			email: "ad@email.com",
		},
		level: "admin",
	}
	sendNotify(&ad)
}

func sendNotify(n notifier){
	n.notify()
}
