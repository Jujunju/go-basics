package main

import (
	"fmt"
)

type MyM interface {
	Sum() int
}

type UserRepo interface {
	Create() *User
}

type I interface {
	Sa()
}

func main() {

	var m MyM
	var u UserRepo

	if as, ok := u.(*User); ok {
		fmt.Println("betul", as.Create())
	}

	mI := myInt(10)
	mS := myStruct{
		X: 10,
		Y: 10,
	}

	mS1 := &myStruct{
		X: 10,
		Y: 10,
	}

	fmt.Println(mS1.test())

	m = mI
	fmt.Println(m.Sum())

	m = &mS
	fmt.Println(m.Sum())

	user1 := &User{
		FullName: "Jujun Junaedi", Username: "jujun_123", Password: "jujun",
	}

	u = user1

	GetInfo(u)

	var n I = &S{
		S: "hello",
	}

	n.Sa()

}

type User struct {
	FullName, Username, Password string
}

type S struct {
	S string
}

type myInt int

func (m myInt) Sum() int {
	return 10
}

type myStruct struct {
	X, Y int
}

func (ms *myStruct) Sum() int {
	return ms.X + ms.Y
}

func (msY *myStruct) test() int {
	return msY.X
}

func (u *User) Create() *User {
	return &User{
		FullName: u.FullName,
		Username: u.Username,
		Password: u.Password,
	}
}

func GetInfo(u UserRepo) {

	u.Create()

	fmt.Println("data user berhasi dibuat")

}

func (s *S) Sa() {}
