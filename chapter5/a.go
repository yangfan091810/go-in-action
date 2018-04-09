package main

import "fmt"

//用户自定义的类型 user
type user struct {
	name  string
	email string
}

//使用值接受者实现一个方法
func (u user) notify() {
	fmt.Printf("userinfo name %s, email %s\n", u.name, u.email)
}

//使用指针接受者实现一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{
		name:  "bill",
		email: "bill@qq.com",
	}
	bill.notify()
	//user类型的值可以调用指针接收者实现的方法
	bill.changeEmail("newbill@qq.com")
	bill.notify()

	lisa := &user{
		name:  "lisa",
		email: "lisa@qq.com",
	}
	lisa.notify()
	lisa.changeEmail("newlisa@qq.com")
	lisa.notify()
}
