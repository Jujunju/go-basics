package main

import "fmt"

func main() {

	val := []int{10, 20, 30, 40, 50, 60}

	getI := val[2:3]

	fmt.Println(getI)

}