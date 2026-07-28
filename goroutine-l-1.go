package main

import (
	"fmt"
	"time"
)

func main() {
	go Si()
	go Si()
	
	fmt.Println("End")
	time.Sleep(500 * time.Millisecond)
}

func Si() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hi2 - ", i)
		time.Sleep(100 * time.Millisecond)
	}
}
