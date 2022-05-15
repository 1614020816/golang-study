package main

import "fmt"

func main()  {
	var a = 10
	var b = 20

	swap(&a, &b)

	fmt.Println("a = ", a, "b = ", b)

	var p *int

	p = &a

	fmt.Println("&a = ", &a)
	fmt.Println("p = ", p)

	// 二级指针
	var pp **int

	pp = &p

	fmt.Println("&p = ", &p)
	fmt.Println("pp = ", pp)
}

//func swap (a int, b int) {
//	var temp int
//	temp = a
//	a = b
//	b = temp
//}

func swap (pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}
