package main

import "fmt"

type User struct {
	FullName string
	Username string
	Password string
}

var arrays = [2]string{"jujun", "junaedi"}

var slices = []string{"lenovo", "asus"}

var maps = map[string]string{
	"user1": arrays[0],
	"tech":  slices[0],
}

func main() {

	user := &User{
		FullName: arrays[0] + arrays[1],
		Username: arrays[0] + "_" + arrays[1],
		Password: arrays[0],
	}
	fmt.Println(user)

}

func rArray() *[2]string {
	arr := [2]string{"alif", "ridwan"}
	return &arr
}

func rSlice() *[]string {
	arr := []string{"thinkpad", "rog"}
	return &arr
}

func rMaps() *map[string]string {
	return &map[string]string{
		"user1": arrays[1],
		"tech":  slices[1],
	}
}

func rStruct() *User {
	return &User{
		FullName: arrays[0] + arrays[1],
		Username: arrays[0] + "_" + arrays[1],
		Password: arrays[0],
	}
}

func (u *User) Create() *User {
	return &User{
		FullName: arrays[0] + arrays[1],
		Username: arrays[0] + "_" + arrays[1],
		Password: arrays[0],
	}
}
