package main

import "fmt"

type W struct {
	n int
	ns []int
	nsi [2]int
}

var maps = map[string]int{"first": 1}

func main() {

	st := W{n: 30, ns: []int{20}, nsi: [2]int{10}}

	myM(maps)
	mys(st)
	
	fmt.Println(maps)
	fmt.Println(st)
}

func myM(m map[string]int) {
	m["first"] = 10
}

func mys(ms W) {
	ms.n = 10
	ms.ns[0] = 10
	ms.nsi[0] = 20
	ms.nsi[1] = 20
}
