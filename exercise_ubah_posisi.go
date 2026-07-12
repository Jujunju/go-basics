package main

import (
	"fmt"
)

func main() {

	balikPosisi("jujun")

}

func balikPosisi(s string) {

	vs := []byte{s[0], s[1], s[2], s[3], s[4]}

	vs[0], vs[1], vs[2], vs[3], vs[4] = vs[4], vs[3], vs[2], vs[1], vs[0]
	
	fmt.Println(vs)

}
