package main

import "fmt"

type T struct {
	User
}

type User struct {
	Name string
	Age int
}

func (u *User) SetName1() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)

	fmt.Println(t.Name)
	t.User.SetName1()
	fmt.Println(t)
}