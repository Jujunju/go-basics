package main

import (
	"io"
	"os"
	"strings"
)

type rot13reader struct {
	reader io.Reader
}

func main() {

	my := strings.NewReader("Zl Anzr Vf Whwha")

	r := rot13reader{my}

	io.Copy(os.Stdout, &r)

}

func (r rot13reader) Read(b []byte) (int, error) {

	n, e := r.reader.Read(b)

	for i := 0; i < n; i++ {
		c := b[i]

		if c >= 'A' && c <= 'Z' {
			b[i] = 'A' + (c-'A'+13)%26
		}

		if c >= 'a' && c <= 'z' {
			b[i] = 'a' + (c-'a'+13)%26
		}

	}
	return n, e

}
