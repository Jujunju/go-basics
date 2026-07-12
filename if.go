// package main

// import (
// 	"fmt"
// 	"math"
// )

// func fullName(firstName, lastName string) string {
// 	return fmt.Sprintln(firstName, lastName)
// }

// func sqrt() float64 {
// 	return math.Sqrt(64)
// }

// func pow(x, y, lim float64) float64 {
// 	if v := math.Pow(x, y); v < lim {
// 		return v
// 	} else {
// 		fmt.Printf("%g >= %g\n", v, lim)
// 	}

// 	return lim
// }

// func Sqrt(x float64) float64 {
// 	z := 1.0

// 	for i := 1; i < 10; i++ {
// 		z -= (z*z - x) / (2*z)
// 		fmt.Println("tebakan saat ini", z)
// 	}

// 	return z
// }

// func main() {
// 	fmt.Println(fullName("jujun", "junaedi"))
// 	fmt.Println(sqrt())
// 	fmt.Println(
// 		pow(2, 3, 50),
// 	)
// 	fmt.Println(
// 		Sqrt(2),
// 	)
// }