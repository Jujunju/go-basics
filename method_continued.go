// package main

// import "fmt"

// type myStr string
// type myInt2 []int

// func main() {
// 	var m myStr = "Jujun"
// 	var mm myInt2 = []int{
// 		10, 20, 30, 40,
// 	}

// 	fmt.Println(m.say())
// 	fmt.Println(mm.findMax())

// 	s := struct {
// 		X, Y int
// 	}{
// 		X: 10,
// 		Y: 10,
// 	}
// 	fmt.Println(s)
// }

// func (m myStr) say() myStr {
// 	return m
// }

// func (i myInt2) findMax() []int {

// 	n := make([]int, 0, 10)

// 	m := i[0]

// 	for _, v := range i {
// 		if v > m {
// 			m = v
// 		}
// 	}

// 	n = append(n, m)

// 	return n
// }
