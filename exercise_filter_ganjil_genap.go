package main

import "fmt"

func main() {

	val := []int{10, 8, 46, 28, 17, 55, 5, 4, 32, 7, 9}

	genap, ganjil := filter(val)

	fmt.Println("Genap =", genap)
	fmt.Println("Ganjil =", ganjil)

}

func filter(v []int) ([]int, []int) {

	genap := make([]int, 0, 11)
	ganjil := make([]int, 0, 11)

	for i := 0; i < len(v); i++ {

		if v[i]%2 == 0 {
			genap = append(genap, v[i])
		}

		if v[i]%2 == 1 {
			ganjil = append(ganjil, v[i])
		}
		
	}

	return genap, ganjil

}


