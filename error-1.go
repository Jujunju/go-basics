package main

import (
	"fmt"
	"log"
)

type Err struct {
	Msg string
}

type User struct {
	Name string
}

type NewUser interface {
	Create() (*User, error)
}

func main() {

	s, e := TrigErr("jujun")

	if e != nil {
		log.Fatal(e.Error())
	}

	st := fmt.Sprintf("success : %s", s)
	
	fmt.Println(st)
	
	n, e1 := Create("ju")
	
	if e1 != nil {
		log.Fatal( e1.Error())
	}
	
	fmt.Println(n)
}

func (u *Err) Error() string {
	return fmt.Sprintf("Error : %s", u.Msg)
}


func TrigErr(name string) (string, error) {

	if name == "" {
		return "", hdErr("nama tidak boleh kosong")
	}

	return "halo " + name, nil
}

func hdErr(msg string) *Err {
	return &Err{Msg: msg}
}

func Create(nameC string) (*User, error) {
	if len(nameC) < 3 {
		return &User{}, hdErr(fmt.Sprintf("nama tidak boleh kurang dari %d", len(nameC)))
	}

	return &User{Name: nameC}, nil
}