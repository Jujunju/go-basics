package main

import "fmt"

func main() {

	val := func(value int) int {
		return value
	}


	names := func(firstName, lastName string) string {
		return firstName + lastName
	}

	fmt.Println(sum(val))
	filter2(names)

}

func sum(fn func(int) int) int {
	return fn(200) + fn(200)
}

func filter(fn func(string) string) {
	re := fn("halo")

	if re == "hussh" {
		fmt.Println("*****")
	} else {
		fmt.Println(re)
	}
}

func filter2(fn func(string, string) string) string {
	return fn("jujun", "junaedi")
}
