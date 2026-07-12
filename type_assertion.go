// package main

// import "fmt"

// type U struct {
// 	Name string
// }

// type I interface {
// 	testU()
// }

// func main() {

// 	var i I
// 	var si U

// 	si.test2()

// 	i = &U{
// 		Name: "jujun",
// 	}

	
// 	if r, s := i.(*U); s {
// 		fmt.Println("saya ambil nih", r.Name)
// 	}

// }

// func (u *U) testU() {
// 	fmt.Println(u.Name)
// }

// func (u *U) test2() {
// 	fmt.Println(u.Name)
// }