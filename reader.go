package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	t := strings.NewReader("Jujun Junaedi")

	m := make([]byte, 10)

	for {

		s, e := t.Read(m)

		fmt.Printf("s = %v m = %v err = %v\n", s, m, e)

		if e == io.EOF {
			break
		}

	}

}