package main

import "fmt"

type Y struct {
	data int
	next *Y
}

func main() {

	t := &Y{data: 10}

	appendY(t, 20)
	appendY(t, 30)

	cur := t

	for cur != nil {
		fmt.Printf("%d -> ", cur.data)
		cur = cur.next
	}

	fmt.Print("nil")

}

func appendY(y *Y, add int) {

	m := &Y{data: add}

	cur := y

	for cur.next != nil {
		cur = cur.next
	}

	cur.next = m

}