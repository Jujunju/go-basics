package main

func main() {
	getResult(10)
}

func getResult(v int) int {

	r := 0

	for i := 1; i < v; i++ {
		r += i
	}

	return r

}
