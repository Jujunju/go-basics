// package main

// import "fmt"

// type Lum[T any] struct {
// 	Next *Lum[T]
// 	V    T
// }

// func main() {

// 	y := &Lum[int]{V: 10}

// 	y.cAdd(20)
// 	y.cShow()

// }

// func (l *Lum[T]) cAdd(i T) {

// 	mw := &Lum[T]{V: i}

// 	c := l

// 	for c.Next != nil {
// 		c = c.Next
// 	}

// 	c.Next = mw

// }

// func (l *Lum[T]) cShow() {
// 	c := l

// 	for c != nil {
// 		fmt.Printf("[%v] -> ", c.V)
// 		c = c.Next
// 	}

// 	fmt.Println("End")
// }