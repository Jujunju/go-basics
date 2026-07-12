package main

import (
	"fmt"
	"time"
)

type User struct {
	ID        uint
	FullName  string
	Username  string
	Password  string
	Level     string
	CreatedAt time.Time
}

func main() {

	user1 := User{1, "Jujun Junaedi", "jujun_junaedi", "jujun123", "visitor", time.Now()}

	user2 := User{
		ID: 2,
		FullName: "Aldi Sanjaya",
		Username: "aldi_sanjaya",
		Password: "aldi123",
		Level: "admin",
		CreatedAt: time.Now(),
	}

	user2.Level = "visitor"

	user3 := &User{}
	user3.ID = 3

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)

}
