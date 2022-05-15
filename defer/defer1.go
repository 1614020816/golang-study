package main

import "fmt"

func main() {
	defer fmt.Println("main end")

	var a int = 100
	var b int = 200

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
