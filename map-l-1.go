package main

import "fmt"

var tem map[string]int = map[string]int{
	"first": 10,
}

func main() {

	s, si := tem["firs"]

	if si {
		fmt.Println(s)
	} else {
		fmt.Println(s)
	}

}
