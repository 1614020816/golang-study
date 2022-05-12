package main

import "fmt"

// 返回单返回值
func foo1(a string, b int) int {
	fmt.Println("----------Foo1----------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

// 返回多个返回值，匿名
func foo2(a string, b int) (int, int){
	fmt.Println("----------Foo2----------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 333, 444
}

// 返回多个返回值，有形参名称的,同一类型的返回值还可以写成 (r1, r2 int)
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----------Foo3----------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// r1 r2 属于foo3的形参，初始化默认值为0
	// r1 r2 作用域空间是整个foo3函数的{}区间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	r1 = 1000
	r2 = 2000

	return
}

func main()  {
	c := foo1("abc", 123)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("abc", 123)
	fmt.Println("rrt1 = ",ret1, "ret2 = ", ret2)

	ret3, ret4 := foo3("abc", 200)
	fmt.Println("ret3 = ",ret3, "ret4 = ", ret4)
}
