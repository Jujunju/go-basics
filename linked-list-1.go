package main

import "fmt"

type User struct {
	data int
	next *User
}

func main() {
	d1 := &User{data: 10}
	d2 := &User{data: 20}
	d3 := &User{data: 30}
	d4 := &User{data: 40}
	d5 := &User{data: 50}

	d1.next = d2
	d2.next = d3
	d3.next = d4
	d4.next = d5

	curr := d1

	for curr != nil {
		fmt.Printf("%d -> ", curr.data)
		curr = curr.next
	}

	fmt.Println("nil")

}