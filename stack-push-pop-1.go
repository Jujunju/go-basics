package main

import "fmt"

type Stack []int

func main() {

	var ps = &Stack{}
	ps.Push(10)
	ps.Push(20)
	ps.Push(30)
	ps.Push(40)

	fmt.Println(ps.Pop())
	fmt.Println(ps)
	

}

func (s *Stack) Push(d int) {
	*s = append(*s, d)
}

func (s *Stack) Pop() int {

	if len(*s) == 0 {
		return -1
	}

	index := len(*s) - 1
	elemen := (*s)[index]
	*s = (*s)[:index]
	return elemen

}
