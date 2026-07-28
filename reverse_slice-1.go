package main

import (
	"fmt"
)

func main() {

	val := []int{10, 20, 30, 40, 50, 60}

	reverseS(&val)

	fmt.Println(val)

	val2 := []string{"saya", "belajar", "php", "dan", "php"}

	searchR(&val2, "php", "golang")

	fmt.Println(val2)

}

func reverseS(v *[]int) {
	for i := 0; i < len(*v)-1; i++ {
		for j := 0; j < len(*v)-i-1; j++ {
			if (*v)[j] < (*v)[j+1] {
				(*v)[j], (*v)[j+1] = (*v)[j+1], (*v)[j]
			}
		}
	}
}

func searchR(vv *[]string, tar, gan string) {
	for i := 0; i < len(*vv); i++ {
		if (*vv)[i] == tar {
			(*vv)[i] = gan
		}
	}
}
