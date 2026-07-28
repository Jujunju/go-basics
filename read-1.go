package main

import (
	"fmt"
	"io"
)

type MyReader struct {
	data string
	pos  int
}

func main() {

	n := &MyReader{data: "jujun"}

	mt := make([]byte, 3)

	for {
		n, e := n.Read(mt)

		if e != nil {
			fmt.Println("selesai baca")
			break
		}

		fmt.Printf("Baca %d byte: %s\n", n, mt[:n])
	}

}

func (r *MyReader) Read(p []byte) (n int, e error) {

	if r.pos >= len(r.data) {
		n = 0
		e = io.EOF
		return
	}

	n = copy(p, r.data[r.pos:])

	r.pos += n

	return n, nil

}