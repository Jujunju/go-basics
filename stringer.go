// package main

// import "fmt"

// type P struct {
// 	Name string
// 	Age int
// }

// func main() {

// 	p1 := P{
// 		Name: "jujun",
// 		Age: 18,
// 	}
// 	p2 := P{
// 		Name: "yono",
// 		Age: 20,
// 	}

// 	p1.String()

// 	fmt.Println(p1)
// 	fmt.Println(p2)

// }

// func (p P) String() string {
// 	return fmt.Sprintf("%v ", p.Name)
// }