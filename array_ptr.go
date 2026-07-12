package main

import "fmt"

var users [3]string = [3]string{"sandira", "ayu", "anli"}
var address [3]string = [3]string{"jl sutomo", "jl mangku", "jl simantri"}

func main() {

	u1 := &users
	a1 := &address

	str := &users[0]

	*str = "jujun"

	a1[0] = "jl jambrud"

	for i, v := range a1 {
		fmt.Println(i, v)
	}

	fmt.Println(users)
	fmt.Println(address)

	fmt.Println(u1)
	fmt.Println(a1)


}