package main

import "fmt"

type Stack []int

func main() {

	var s Stack

	s.Enqueue(10)
	s.Enqueue(20)
	s.Enqueue(30)

	fmt.Println(s)
	fmt.Println(s.Dequeue())

}

func (s *Stack) Enqueue(v int) {

	*s = append(*s, v)

}

func (s *Stack) Dequeue() int {
	if len(*s) == 0 {
		return -1
	}

	elemen := (*s)[0]

	*s = (*s)[1:]

	return elemen
}