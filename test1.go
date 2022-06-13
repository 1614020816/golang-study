package main

import "fmt"

// const 和 iota
func main() {
	const len int = 100
	fmt.Println("len = ", len)

	const (
		BEIJING = 10 * iota
		SHANGHAI
		SHENZHEN
		GUANGZHOU
	)

	const (
		A = 10
		B = 20
		C = 30
	)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)
	fmt.Println("GUANGZHOU = ", GUANGZHOU)

	const (
		a, b = iota + 1, iota + 2   // iota = 0, a = 1, b = 2
		c, d                        // iota = 1, c = 2, d = 3
		e, f                        // iota = 2, e = 3, f = 4

		g, h = iota * 2, iota * 3   // iota = 3, g = 6, h = 8
		i, j                        // iota = 4, i = 8, j = 12
	)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

	// iota 只能配合 const() 使用, 只有在 const 中进行累加效果
}