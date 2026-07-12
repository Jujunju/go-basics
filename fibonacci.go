package main

import "fmt"

func main() {

	s := sum()

	for i := 0; i <= 10; i++ {
		fmt.Println(s())

	}

}

func sum() func() int {

	r1 := 0
	r2 := 1

	return func() int {
		t := r1

		r1, r2 = r2, r1 + r2

		return t
	}

}
