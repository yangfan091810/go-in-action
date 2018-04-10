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
	fmt.Printf("user name %s email %s \n", u.name, u.email)
}

type admin struct {
	name	string
	email	string
}

func (a *admin) notify() {
	fmt.Printf("admin name %s, email %s \n", a.name, a.email)
}

func main(){
	lili := &user{
			name: "lili",
			email: "lili@lili.com",
		}
	sendNotify(lili)

	haha := &admin{
		name: "haha",
		email: "haha@haha.com",
	}
	sendNotify(haha)
}

func sendNotify(n notifier){
	n.notify()
}
