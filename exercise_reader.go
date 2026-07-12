package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func main() {

	reader.Validate(MyReader{})

}

func (m MyReader) Read(b []byte) (int, error) {

	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}

	return len(b), nil

}
