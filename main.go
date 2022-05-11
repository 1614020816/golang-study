package main

import "fmt"

func main() {
	// 没有赋初始值 默认为0
	var start int
	fmt.Println("start = ", start)

	var a int = 100
	fmt.Println("a = ", a)

	var b = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	// 这种命名方式只能用于局部变量声明
	c := 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var d, e int = 100, 200
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)

	var f, g = 100, "name"
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	// 多变量声明方式
	var (
		i int = 100
		j bool = false
	)

	fmt.Println("i = ", i)
	fmt.Printf("type of i = %T\n", i)
	fmt.Println("j = ", j)
	fmt.Printf("type of j = %T\n", j)
}
