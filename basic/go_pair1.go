package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

//具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book.")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book.")
}

func main() {
	b := &Book{}

	var r Reader
	r = b

	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriteBook()
}
