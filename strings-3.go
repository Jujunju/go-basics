package main

import (
	"fmt"
	"strings"
)

func main() {
	t := strings.Contains("jujun", "j")

	vv := "jujunjunaedi"

	vv2 := strings.Clone(vv)

	fmt.Println(vv2 == vv)

	ni := strings.Split(vv, " ")

	vvi := []string{"jujun", "junaedi"}

	gt := strings.Join(vvi, ", ")

	rp := "jujun suka ngoding, dan jujun juga suka makan dan minum"

	fmt.Printf("%d\n", strings.Count(rp, "s"))

	ni2 := filRp(rp)

	rpAll := strings.ReplaceAll(rp, "suka", "mau")

	fmt.Println(rpAll)
	fmt.Println(ni2)

	rp = strings.Replace(rp, "suka", "mau", ni2)

	hf := strings.HasPrefix("ju-jun", "ju")
	hs := strings.HasSuffix("ju-jun", "un")

	fmt.Printf("bool = %t\n", hf)
	fmt.Printf("bool = %t\n", hs)
	fmt.Println(t)
	fmt.Printf("%q\n", ni)
	fmt.Println(gt)
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Println(strings.Trim(" -jujun", " -"))
	fmt.Println(strings.TrimSpace(" -jujun"))
	fmt.Println(rp)
	fmt.Printf("%d\n", strings.Index(rp, "mau"))
	fmt.Printf("%d\n", strings.Index(rp, "ngoding"))

}


func filRp(v any) int {
	result := make([]string, 0, 100)
	if s, e := v.(string); e {
		element := strings.Split(s, " ")

		for i := 0; i < len(element); i++ {
			if element[i] == "suka" {
				result = append(result, element[i])
			}
		}
	}
	return len(result)
}
