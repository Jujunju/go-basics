package main

import (
	"fmt"
)

type User struct {
	Name, Address string
}

var m3 map[string]string
var s [3]string
var s3 []string

func main() {

	m1 := make(map[int]int, 10)
	m2 := make(map[string]*User, 10)
	m3 := make(map[string]string, 10)

	myMaps := map[string]string{
		"test1": "..",
		"test2": "...",
		"test3": "....",
	}

	myMaps["test1"] = "Iphone"

	delete(myMaps, "test3")

	if v, ok := myMaps["test3"]; ok {
		fmt.Println("data ada", v)
	} else {
		fmt.Println("data sudah tidak ada", v)
	}

	m3["n1"] = "yana"
	m3["n2"] = "sana"
	m3["n3"] = "bona"

	m3["n3"] = "john"

	s[0] = "jujun"

	m1[0] = 100

	m2["u1"] = &User{
		Name:    "jujun",
		Address: "jl selamat",
	}
	m2["u2"] = &User{
		Name:    "bakri",
		Address: "jl sinar mas",
	}

	m2["u1"].Address = "jl haji"

	s3 = append(s3, "jujun")

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(s)
	fmt.Println(s3)
	fmt.Println(m3)
	fmt.Println(myMaps)

}
