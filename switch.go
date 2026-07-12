package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	var name string = "jujun"

	switch name {
	case "jujun":
		fmt.Println("Ya betul")
	default:
		fmt.Println("Tidak valid")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("mac")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("os", os)
	}

	fmt.Printf("%T %v", time.Now().Hour(), time.Now().Minute())

	today := time.Now().Weekday()
	switch today {
	case time.Sunday:
		fmt.Println("minggu")
	case time.Monday:
		fmt.Println("senin")
	case time.Tuesday:
		fmt.Println("selasa")
	case time.Wednesday:
		fmt.Println("rabu")
	case time.Thursday:
		fmt.Println("kamis")
	case time.Friday:
		fmt.Println("jumat")
	case time.Saturday:
		fmt.Println("sabtu")
	default:
		fmt.Println("?")
	}

	fmt.Println(today)

	c := 10
	switch {
	case c < 10:
		fmt.Println("betul")
	case c > 20:
		fmt.Println("salah")
	default:
		fmt.Println("all")

	}
}
