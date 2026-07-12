package main

import "fmt"

type User struct {
	Name, Username string
}

type UserRepository interface {
	Create() *User
}

type NewUser struct {}

func main() {

	repo := NewURepository()

	newU := repo.Create()

	fmt.Println(newU)

}

func (repo *NewUser) Create() *User {
	return &User{Name: "Jujun", Username: "jujun_123"}
}

func NewURepository() UserRepository {
	return &NewUser{}
}
