package main

import "fmt"

var slices = []string{"wortel", "pisang", "anggur"}

func main() {

	fmt.Printf("len=%d cap=%d %v\n", len(slices), cap(slices), slices)

}