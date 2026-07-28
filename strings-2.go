package main

import (
	"fmt"
	"strings"
)

func main() {

	n := strings.Contains("jujun", "n")

	ni := strings.Fields("jujun junaedi ini")

	c := strings.Count("wi", "i")

	fmt.Printf("%t", n)
	fmt.Printf("\n%d", len(ni))
	fmt.Printf("\n%d", c)

}