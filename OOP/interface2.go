package main

import "fmt"

type Book2 struct {
	name string
}

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	fmt.Println("arg = ", arg)

	// 使用interface的'类型断言'机制，推断当前底层数据类型是什么
	value, ok := arg.(string)

	if !ok {
		fmt.Println("arg type is not string")
	} else {
		fmt.Println("arg type is string, value = ", value)
		fmt.Printf("value type is %T\n", value)
	}
}

func main() {
	book := Book2{"Golang"}

	myFunc(book)

	myFunc(100)

	myFunc("abc")

	myFunc(3.141)
}