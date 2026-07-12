// package main

// import (
// 	"fmt"
// )

// func main() {

// 	n := cleanDuplicateContact("08556677", "0776655", "08556677", "08447766", "08669933", "086754657", "08447766")

// 	fmt.Println("bersih = ", n)
// }

// func cleanDuplicateContact(l ...string) []string {

// 	y := make(map[string]bool, 100)
// 	yv := make([]string, 0, 100)

// 	for _, v := range l {
// 		if !y[v] {
// 			y[v] = true
// 			yv = append(yv, v)
// 		}
// 	}

// 	return yv
// }
