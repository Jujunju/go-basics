package main

import "fmt"

func main() {
	val := []int{77, 65, 100, 87, 34, 55}
	bubbleSortMax(&val)

	fmt.Println(val)

	val2 := []int{77, 65, 100, 87, 34, 55}
	bubbleSortMin(&val2)

	fmt.Println(val2)
}

func bubbleSortMax(v *[]int) {
	n := len(*v)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*v)[j] > (*v)[j+1] {
				(*v)[j], (*v)[j+1] = (*v)[j+1], (*v)[j]
			}
		}
	}
}

func bubbleSortMin(vl *[]int) {
	n := len(*vl)

	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - i - 1; j++ { 
			if (*vl)[j] < (*vl)[j+1] {
				(*vl)[j], (*vl)[j+1] = (*vl)[j+1], (*vl)[j]
			}
		}
	}
}