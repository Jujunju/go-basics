package main

import (
	"errors"
	"fmt"
)

var (
	NotFoundError       = errors.New("Error : NotFound")
	BadRequestError     = errors.New("Error : BadRequest")
	InternalServerError = errors.New("Error : InternalServer")
)

func main() {

	n, e := SayHello("jujun", false)

	if errors.Is(e, NotFoundError) {
		fmt.Println(e.Error())
		return
	}
	if errors.Is(e, BadRequestError) {
		fmt.Println(e.Error())
		return
	}

	fmt.Println(n)

}

func SayHello(name string, errSrv bool) (n string, err error) {

	if name == "" {
		n = ""
		err = BadRequestError
		return n, err
	}

	if name != "jujun" {
		n = ""
		err = NotFoundError
		return n, err
	}

	if errServer := errSrv; errServer {
		n = ""
		err = InternalServerError
		return n, err
	}

	na := struct {
		Name string
	}{
		Name: "jujun",
	}

	return na.Name, nil

}
