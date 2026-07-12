// package main

// import "fmt"

// type Product struct {
// 	Name, Category string
// 	price          uint64
// }

// func main() {

// 	s := sayHello("halo")
// 	m := getInfos()

// 	fmt.Println(s)
// 	fmt.Println(m)

// }

// func sayHello(say string) func() string {
// 	names := "jujun"
// 	return func() string {
// 		return say + names
// 	}
// }

// func getInfos() func() *Product {
// 	nameP := "Infinix smart 4"
// 	cateP := "Elektronik"
// 	prP := 1000000
// 	myFunc := func() *Product {
// 		return &Product{
// 			Name:     nameP,
// 			Category: cateP,
// 			price:    uint64(prP),
// 		}
// 	}

// 	return myFunc
// }
