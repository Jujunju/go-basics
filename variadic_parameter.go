package main

import (
	"fmt"
	"strings"
)

var names = []string{"jujun", "alif", "yono"}

func main() {

	example([]string{"jujun", "alif", "yono"}...)
	fmt.Println(example(names...))

	fmt.Println(maxVal(10, 20, 30, 40, 50))
}

func example(name ...string) string {
	myNames := strings.Join(name, " - ")

	return myNames
}

func maxVal(val ...int) int {

	if len(val) == 0 {
		return 0
	}

	max := val[0]

	for _, v := range val {
		if v > max {
			max = v
		}
	}

	return max
	
}
