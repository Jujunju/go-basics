// package main

// import "fmt"

// type MyI interface {
// 	Me() string
// }

// var n interface{}

// func main() {

// 	var mI MyI

// 	n = mI
// 	n = true

// 	if r, s := n.(MyI); s {
// 		fmt.Println(r.Me())
// 	}

// 	switch v:=  n.(type) {
// 	case int:
// 		fmt.Println("ini adalah integer", v)
// 	case string:
// 		fmt.Println("ini adalah string", v)
// 	case bool:
// 		fmt.Println("ini adalah boolean", v)
// 	default:
// 		fmt.Println("tidak dikenali", v)
// 	}
// }

// func Me() string {
// 	return "hi"
// }
