package main

import "fmt"

func main() {
	n, l, e := myApp("Magic Chess", 12, false)

	fmt.Println(n, l, e)
}

func myApp(name string, layout int, exists bool) (name1 string, layout1 int, exists1 bool) {

	name1 = name
	layout1 = layout
	exists1 = exists

	return name1, layout1, exists1

}
