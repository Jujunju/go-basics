package main

import (
	"fmt"
	"time"
)

type AuthResponse struct {
	ID        uint
	Username  string
	Group     string
	Token     string
	CreatedAt time.Time
}

func (a *AuthResponse) getUsername() *string {
	return &a.Username
}

func main() {

	u1 := &AuthResponse{
		ID:        1,
		Username:  "jujun_junaedi",
		Group:     "admin",
		Token:     "92638dsgd927&%(@#flfvsfsm",
		CreatedAt: time.Now(),
	}

	u2 := &AuthResponse{}

	fmt.Println(u1.getUsername())
	fmt.Println(u1)
	fmt.Println(u2)

}
