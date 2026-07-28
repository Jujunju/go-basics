package main

import "fmt"

type User struct {
	Username string
}

func main() {

	n := C()
	for i:=0;i<10;i++ {
		fmt.Println(n())
	}
	
	m := P(func() string {
		return "JUJUN JUNAEDI"
	})
	fmt.Println(m().Username)

}

func C() func()int {
	l := 0

	return func() int {
		l += 1
		return l
	}
	
}

func P(v func() string) func()*User {
	n := v()
	return func() *User {
		return &User{Username: n}
	}
}
