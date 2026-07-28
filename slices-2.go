package main

import (
	"cmp"
	"fmt"
	"slices"
)

type User struct {
	Name string
	Age  int
}

func main() {

	for i, v := range slices.Backward([]string{"jujun", "lipa", "manda"}) {
		fmt.Printf("index : %d element : %s\n", i, v)
	}

	n, b := slices.BinarySearch([]string{"jujun", "lipa", "manda"}, "mmii")
	fmt.Printf("n : %d b : %t\n", n, b)

	fmt.Printf("%d\n", slices.Compare([]string{"jujun", "lipa", "manda"}, []string{"jujun", "lipa", "manda"}))

	nm := []string{"jujun", "lipa", "manda"}

	nm = slices.Insert(nm, len(nm), "jamilah", "yono", "apim")

	nmy := slices.Concat(nm, []string{"yanli"})

	fmt.Printf("\ninsert %q ", nm)
	fmt.Printf("\nconcat %q ", nmy)
	fmt.Printf("\n%d ", slices.Index(nm, "lipa"))
	nms := slices.Delete(nm, 1, 2)
	fmt.Printf("%q ", nms)

	fmt.Printf("\nMax %d\n", slices.Max([]int{10, 20, 30, 40}))
	fmt.Printf("Min %d\n", slices.Min([]int{10, 20, 30, 40}))

	slicesM := []User{
		{
			Name: "jujun",
			Age:  18,
		},
		{
			Name: "yono",
			Age:  22,
		},
		{
			Name: "afud",
			Age:  30,
		},
	}

	vvi := []string{"agung", "sami", "fami"}

	slices.Reverse(slicesM)
	slices.SortFunc(slicesM, func(a, b User) int {
		return cmp.Compare(a.Age, b.Age)
	})
	slices.Reverse(vvi)

	fmt.Printf("sort %q\n", slicesM)

	fd := slices.MaxFunc(slicesM, func(a, b User) int {
		return cmp.Compare(a.Age, b.Age)
	})

	fmt.Println(fd)

	fmi := slices.Replace(slicesM, 1, 2, []User{
		{
			Name: "agus",
			Age:  22,
		},
	}[0])
	fm := slices.Repeat(slicesM, 2)

	fmt.Println(fm)
	fmt.Println(fmi)

}
