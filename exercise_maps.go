package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	hasilHitung := make(map[string]int)

	kumpulanKata := strings.Fields(s)

	for _, kata := range kumpulanKata {
		hasilHitung[kata]++
	}

	return hasilHitung

}

func find(vName ...string) ([]string, []string) {

	resultNames := make([]string, 0, 100)
	cName := make([]string, 0, 100)

	for _, v := range vName {
		switch {
		case vName[0] == v:
			resultNames = append(resultNames, v)
		default:
			cName = append(cName, v)
		}
	}

	return resultNames, cName
}

func main() {

	vv, vv2 := find("jujun", "alif", "yono", "bakri")

	fmt.Println(vv)
	fmt.Println(vv2)

	wc.Test(WordCount)
}
