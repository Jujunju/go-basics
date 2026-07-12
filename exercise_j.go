package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	if h, e := hitungHargaPasti(true); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(h)
	}

	

}

func hitungHargaPasti(i any) (float64, error) {
	switch s := i.(type) {
	case int:
		return float64(s), nil
	case float64:
		return s, nil
	case string:
		n, _ := strconv.ParseFloat(s, 64)
		return n, nil
	default:
		return 0, errors.New("Tipe data tidak didukung")
	}
}