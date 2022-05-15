/*
* defer 和 return 的执行顺序
*/

package main

import "fmt"

// 执行顺序是 defer 在 return 之后
func main() {
	returnAndDefer()
}

// 在写法上保持defer 在 return 之前
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func deferFunc() int {
	fmt.Println("defer func called...")

	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")

	return 0
}