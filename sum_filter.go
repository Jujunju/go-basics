// package main

// import "fmt"

// func main() {

// 	n := hitungTotal([]int{30000, 75000, 55000})

// 	fmt.Println(n)

// }

// func hitungTotal(v []int) int64 {

// 	result := 0

// 	for i := 0; i < len(v); i++ {
// 		result += v[i]
// 	}

// 	if result > 100000 {
// 		result -= 10000
// 	}

// 	return int64(result)

// }