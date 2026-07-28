package main

import (
	"fmt"
	"strconv"
)

func main() {

	n, e := strconv.Atoi("5")
	
	if e != nil {
		fmt.Println(e)
	}
	
	ni:= strconv.Itoa(5)
	
	st, er := strconv.ParseFloat("5555", 32)

	if er != nil {
		fmt.Println(e)
	}

	b, er2 := strconv.ParseBool("hi")

	if er2 != nil {
		fmt.Println(e)
	}

	fmt.Println(n)
	fmt.Println(ni)
	fmt.Println(st)
	fmt.Println(b)

}