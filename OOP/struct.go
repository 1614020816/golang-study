package main

import "fmt"

type Book struct {
	name string
	auth string
}

func changeBook(book Book){
	// 值传递
	book.name = "JS"
}

func changeBook2(book *Book) {
	// 指针传递
	book.name = "TS"
}

func main() {
	var book1 Book
	book1.name = "Golang"
	book1.auth = "zhangsan"

	fmt.Println("book1 = ", book1)
	fmt.Printf("type of book1 %T\n", book1)

	changeBook(book1)
	fmt.Println("book1 = ", book1)

	changeBook2(&book1)
	fmt.Println("book1 = ", book1)
}