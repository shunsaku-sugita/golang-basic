package main

import "fmt"

type User struct {
	Name string
	Age int
	// X, Y int // 複数のフィールドをまとめて定義することもできる
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	user2.Age = 20
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := User{"user4", 40} // フィールド名を省略する場合は、構造体のフィールド順に値を指定する
	fmt.Println(user4)

	user5 := User{Name: "user5"}
	fmt.Println(user5)

	user6 := new(User)
	fmt.Println(user6)

	user7 := &User{}
	fmt.Println(user7)

	UpdateUser(user1)
	fmt.Println(user1)

	UpdateUser2(user7)
	fmt.Println(*user7)

}