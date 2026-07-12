package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Error : %v\n", float64(e))
}

func main() {

	i, e := Sqrt(-1)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Println(i)


}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z, nil
}
