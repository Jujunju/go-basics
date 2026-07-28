package main

import (
	"errors"
	"fmt"
)

type NewUser interface {
	Create(nama_lengkap, username string) error
	FindByUsername(username string) (*User, error)
}

type User struct {
	NamaLengkap, Username string
}

type Error struct {
	code int
	msg  string
}

var (
	NotFoundError       = errors.New("Error : NotFound")
	BadRequestError     = errors.New("Error : BadRequest")
	InternalServerError = errors.New("Error : InternalServer")
)

func main() {

	user := &User{}
	if e := user.Create("jujun junaedi", "jujun_123"); e != nil {
		var s *Error
		if errors.As(e, &s) {
			fmt.Println(e.Error())
		}
	}

	fmt.Println(user)

	n, e := user.FindByUsername("jujun_123")
	
	if e != nil {
		var si *Error
		if errors.As(e, &si) {
			fmt.Println(e.Error())
		}
	}

	fmt.Println(n)

}

func (e *Error) Error() string {
	return fmt.Sprintf("code = %d\nmsg = %s", e.code, e.msg)
}

func HandleError(code int, msg string) (e error) {
	e = &Error{code: code, msg: msg}
	return
}

func (u *User) Create(nama_lengkap, username string) (e error) {

	if nama_lengkap == "" || username == "" {
		e = HandleError(400, BadRequestError.Error())
		return
	}

	u.NamaLengkap = nama_lengkap
	u.Username = username

	e = nil
	return

}

func (u *User) FindByUsername(username string) (v *User, e error) {

	if u.Username != username {
		v = nil 
		e = HandleError(404, NotFoundError.Error())
		return
	}

	v = &User{NamaLengkap: u.NamaLengkap, Username: u.Username}
	e = nil

	return

}
