package main

import "fmt"

type User struct {
	X, Y int
}

func main() {

	user := User{
		X: 10,
		Y: 10,
	}

	user2 := &User{
		X: 10,
		Y: 10,
	}

	user.test()
	test2(*user2)

	user.test()

	fmt.Println(user.X)
	fmt.Println(user2.X)

}

func (u *User) test() {
	u.X = 5
}

func test2(u User) {

}