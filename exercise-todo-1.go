package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type NewTodo interface {
	Create() error
	Show() (*DataTodo, error)
}

type DataTodo struct {
	Judul, Isi string
}

type Error struct {
	msg string
}

func main() {

	MainTodoApp()

}

func (e *Error) Error() string {
	return fmt.Sprintf("error = %s\n", e.msg)
}

func HandlerErr(msg string) error {
	return &Error{msg: msg}
}

func MainTodoApp() {
	fmt.Println("")

	fmt.Println("==== SELAMAT DATANG DI TODO APP MINI ====")

	fmt.Println("")

	fmt.Println("Pilih Menu :")

	fmt.Println("1. Buat Todo")
	fmt.Println("2. Lihat Todo")
	fmt.Println("3. Hapus Todo")
	fmt.Println("4. Edit Todo")

	fmt.Println("")

	fmt.Print("Masukkan berupa angka 1 - 4 : ")

	rd := bufio.NewScanner(os.Stdin)

	for rd.Scan() {

		d := rd.Text()
		si := &DataTodo{}

		if d == "1" {

			si.Create()

			break
		}

		if d == "2" {
			fmt.Println(si.Show())
			break
		}

	}
}

func (dt *DataTodo) Create() (e error) {
	fmt.Print("Buat Judul : ")
	rdn := bufio.NewScanner(os.Stdin)

	file, err := os.OpenFile("todo-app.txt", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for rdn.Scan() {
		judul := rdn.Text()

		dt.Judul = judul
		io.Copy(file, strings.NewReader("(" + "Judul : " + dt.Judul + ")"))
		break
	}

	fmt.Print("Isi : ")
	rdi := bufio.NewScanner(os.Stdin)

	for rdi.Scan() {
		isi := rdi.Text()

		dt.Isi = isi
		io.Copy(file, strings.NewReader("\n" + "(" + "Isi : "+dt.Isi + ")"))
		break
	}

	fmt.Println("success")

	e = nil

	return
}

func (dt *DataTodo) Show() (tdoi *DataTodo, e error) {

	file, err := os.Open("todo-app.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	dti := &DataTodo{}

	for sc.Scan() {
		text := sc.Text()
		sp := strings.TrimSpace('\n')

		if  {

		}

		dti.Judul = text

		break
	}

	return dti, nil

}
