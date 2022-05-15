package main

import "fmt"

func main() {
	array1 := []int{1, 2, 3, 4}  // 动态数组，切片 slice

	fmt.Printf("array1 types = %T\n", array1)

	printArray2(array1)

	fmt.Println("=======")

	for _, value := range array1 {
		fmt.Println("value = ", value)
	}

}

func printArray2(array []int) {
	// 引用传递, 指针传递，传递的是数组的地址
	for _, value := range array{
		fmt.Println("value = ", value)
	}

	array[0] = 100
}
