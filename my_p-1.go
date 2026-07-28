package main

import "fmt"

func main() {

	mys := []int{
		30,
	}

	mys2 := map[string]string{"first": "alif"}

	reSl(mys, mys2)
	fmt.Println(mys)
	fmt.Println(mys2)

}

func reSl(s []int, ms map[string]string) {
	s[0] = 10
	ms["first"] = "yomo"
}
