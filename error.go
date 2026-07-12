// package main

// import (
// 	"fmt"
// 	"time"
// )

// type err struct {
// 	W time.Time
// 	Y string
// }

// func main() {

// 	if err := hd(); err != nil {
// 		fmt.Println(err)
// 	}

// }

// func (e *err) Error() string {
// 	return e.Y
// }

// func hd() error {
// 	return &err{time.Now(), "jujun"}
// }
