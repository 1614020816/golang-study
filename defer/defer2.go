/*
* defer 的执行顺序
*/
package main

import "fmt"

func main() {
	defer func1()
	defer func2()
	defer func3()
}

func func1() {
	fmt.Println("func1 called ...")
}

func func2() {
	fmt.Println("func2 called ...")
}

func func3() {
	fmt.Println("func3 called ...")
}