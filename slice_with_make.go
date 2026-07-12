package main

import "fmt"

var temp = [3]string{}

func main() {

	values := make([]string, len(temp), cap(temp))

	values[0] = "Jujun"
	values[1] = "alif"
	values[2] = "nugraha"

	values = append(values, "jujun")


	users := make([]string, 0, 10)

	for _, v := range values {
		users = append(users, v)
	}


	fmt.Println(values)
	fmt.Println(users[0])

}