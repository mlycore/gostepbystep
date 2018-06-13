package main

import "fmt"

type user struct{
	name string
	email string
}

//notify implement a method that called via a pointer of type user.
func (u *user)notify(){
	fmt.Printf("Sending username to %s <%s>\n ",u.name,u.email)
}

//admin present an admin user with privileges.
type admin struct {
	user
	level string
}

func main(){
	//creat a admin user
	ad := admin{
		user :user{
			name :"john smith",
			email:"john@163.com",
		},
		level:"super",
	}

	//we can access inner type's method directly.
	ad.user.notify()

	//the inner type's promoted to the outer type.
	ad.notify()
}