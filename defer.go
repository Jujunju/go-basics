package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		defer fmt.Print(i)
	}

	say := func(say string) string {
		return say
	}("hello")

	fmt.Println(say)

	getErr()

}

func getErr() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered ", r)
		}
	}()
	trigErr(12)
}

func trigErr(v int) {
	if v > 10 {
		fmt.Println("no err")
	} else {
		panic("tidak boleh lebih dari angka 10")
	}

	trigErr(v + 1)
}
