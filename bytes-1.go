package main

import (
	"bytes"
	"fmt"
)

func main() {

	bi := []byte("jujun")

	b := bytes.Clone(bi)

	if n := bytes.Contains([]byte("jujun"), []byte("i")); n {
		fmt.Printf("%t\n", n)
	} else {
		fmt.Printf("%t\n", n)
	}

	if ni := bytes.ContainsAny(bi, "j"); ni {
		fmt.Printf("%t\n", ni)
	} else {
		fmt.Printf("%t\n", ni)
	}
	
	if nii := bytes.HasPrefix([]byte("jamilah"), []byte("j")); nii {
		fmt.Printf("prefix : %t\n", nii)
	} else {
		fmt.Printf("prefix : %t\n", nii)
	}

	if ni1 := bytes.HasSuffix([]byte("jamilah"), []byte("j")); ni1 {
		fmt.Printf("suffix : %t\n", ni1)
	} else {
		fmt.Printf("suffix : %t\n", ni1)
	}

	fmt.Println(b[len(b)-1])
	fmt.Printf("s : %d\n", bytes.Count(bi, []byte("uu")))

}
