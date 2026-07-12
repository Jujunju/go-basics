package main

import (
	"errors"
	"fmt"
)

type User struct {
	Username string
	IsAdmin  bool
}

type Sa struct {
	Location string
}

type C interface{}

var (
	Forbidden    = errors.New("Error: Forbidden")
	Unauthorized = errors.New("Error: Unauthorized")
)

func main() {

	var c C

	n, e := validasiSesi(&User{
		Username: "jujun_ju",
		IsAdmin: false,
	})
	c = &User{
		Username: "jujun_ju",
		IsAdmin: false,
	}

	di := &Sa{}

	n, e := validasiSesi(di)

	fmt.Println(n, e)

}

func validasiSesi(sesi any) (string, error) {

	if r, s := sesi.(*User); s {
		if r.IsAdmin {
			return "selamat datang <username>", nil
		}
		return "", Forbidden
	}

	return "", Unauthorized
}
