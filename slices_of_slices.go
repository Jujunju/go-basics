package main

import "fmt"

func main() {

	papan := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	papan[0][0] = "X"
	papan[0][1] = "X"
	papan[0][2] = "X"

	for i, _ := range papan {
		fmt.Println(papan[i])
	}

}