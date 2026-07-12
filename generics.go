// package main

// import (
// 	"fmt"
// )

// func random[T any](i T) T {
// 	return i
// }

// func random2[T any](i []T) []T {
// 	return i
// }

// func random3[T any](s, y T) {

// 	switch rr := any(s).(type) {
// 	case int:
// 		fmt.Println(rr)
// 	}

// }

// func main() {

// 	x := random(10)
// 	v := random("hello")
// 	y := random(map[string]bool{
// 		"c": true,
// 		"b": false,
// 	})
// 	yi := random2([]int{10, 20, 30})

// 	fmt.Println(x)
// 	fmt.Println(v)
// 	fmt.Println(y)
// 	fmt.Println(yi)

// }
