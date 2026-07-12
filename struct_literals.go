// package main

// import "fmt"

// type Person struct {
// 	Name    string
// 	Address string
// 	City    string
// }

// var (
// 	person1 = &Person{Name: "yono", Address: "Jl Mangku", City: "Bandung"}
// 	person2 = &Person{Name: "asri"}
// 	person3 = &Person{Name: "sila", Address: person1.Address}
// 	person4 = func(p *Person) string {
// 		return p.Name
// 	}(person1)
// )

// var is_ptr *string

// func main() {

// 	is_ptr = &person1.Name

// 	a := 10

// 	b := &a

// 	*b = 10

// 	fmt.Println(person4)

// }