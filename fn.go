package main

import "fmt"

func main() {

	post := test()
	xiy := eRe()

	xiy(0)

	fmt.Println(post(5))

}

func test() func(v int) int {
	s := 0
	return func(v int) int {
		s += v
		return s
	}
}

func eRe() func(int) int {
	return func(x int) int {
		return x
	}
}
