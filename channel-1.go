package main

import "fmt"

var n map[string]int = map[string]int{}

func main() {

	n["s"] = 10

	fmt.Println(n)

	pk := make(chan int)

	go HK(10, pk)

	dt := <- pk

	fmt.Println(dt)

}

func HK(a int, pipa chan int) {
	hasil := a * a

	pipa <- hasil
}