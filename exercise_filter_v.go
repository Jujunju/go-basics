// package main

// import "fmt"

// func main() {

// 	var i = []string{"halo", "kamu", "rusuh", "hebat"}

// 	mb := sensorKata(i)

// 	fmt.Println(mb)

// 	var fnFilter = func(ss []string) []string {

// 		temp := []string{}

// 		for i := 0; i < len(ss); i++ {
// 			if ss[i] != "rusuh" {
// 				temp = append(temp, ss[i])
// 			}
// 		}

// 		return temp
// 	}

// 	fmt.Println(sensorKata2(i, fnFilter))

// }

// func sensorKata(s []string) []string {

// 	result := []string{}

// 	for _, v := range s {
// 		if v != "rusuh" {
// 			result = append(result, v)
// 		}
// 	}

// 	return result
// }

// func sensorKata2(s []string, rf func([]string) []string) []string {
// 	result := rf(s)

// 	return  result
// }
