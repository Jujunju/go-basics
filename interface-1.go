package main

import (
	"fmt"
)

type Bul interface {
	GetName() string
}

type S struct{}

func main() {
	var bul Bul

	bul = &S{}

	In(bul)
}

func (s *S) GetName() string {
	return "Jujun Junaedi"
}

func In(b Bul) {
	if s, e := b.(*S); e {
		fmt.Println(s.GetName())
	}
}

func (s *S) Error() string {
	return fmt.Sprintf("Error : ", s.GetName())
}

