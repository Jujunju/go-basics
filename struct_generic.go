package main

import "fmt"

type User[T any] struct {
	ListMap *User[T]
	V       T
	S       []T
}

func main() {

	u := &User[string]{
		ListMap: &User[string]{},
		V:       "hello",
		S:       []string{"world"},
	}

	fmt.Println(&u.V)

}