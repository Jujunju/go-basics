package main

import (
	"errors"
	"fmt"
)

type Error struct {
	Code    int
	Message string
}

func main() {

	s, e := SayHello("jujun", true)

	var err *Error

	if errors.As(e, &err) {
		fmt.Println(e.Error())
		return
	}

	fmt.Println(s)

}

func (e *Error) Error() string {
	return fmt.Sprintf("code = %d\nmessage = %s\n", e.Code, e.Message)
}

func HandleError(code int, msg string) error {
	return &Error{Code: code, Message: msg}
}

func SayHello(name string, errSrv bool) (n string, err error) {

	if name == "" {
		n = ""
		err = HandleError(400, "Error : BadRequest")
		return n, err
	}

	if name != "jujun" {
		n = ""
		err = HandleError(404, "Error : NotFound")
		return n, err
	}

	if errServer := errSrv; errServer {
		n = ""
		err = HandleError(500, "Error : InternalServer")
		return n, err
	}

	na := struct {
		Name string
	}{
		Name: "jujun",
	}

	return na.Name, nil

}
