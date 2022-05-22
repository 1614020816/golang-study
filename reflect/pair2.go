package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {

}

func (this *Book) ReadBook() {
	fmt.Println("read a book")
}

func (this *Book) WriteBook()  {
	fmt.Println("write a book")
}

func main()  {
	// b: pair<type: Book, value: Book{}地址>
	b := &Book{}

	var r Reader

	// r: pair<type: Book, value: Book{}地址>
	r = b

	r.ReadBook()

	var w Writer

	// r: pair<type: Book, value: Book{}地址>
	w = r.(Writer)  // 此处的w为什么能断言成功？是因为r w的类型一样

	w.WriteBook()
}
