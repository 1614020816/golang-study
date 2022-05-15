package main

import "fmt"

func main()  {
	// 固定长度的数组
	var array1 [10]int
	array2 := [10]int{1, 2, 3, 4}
	array3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	for index, value := range array2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	printArray(array3)

	fmt.Printf("type of array1 = %T\n", array1)
	fmt.Printf("type of array2 = %T\n", array2)
	fmt.Printf("type of array3 = %T\n", array3)

	fmt.Println("=========")

	// 调用printArray修改了数组元素后 没有影响原来数组的值
	for _, value := range array3 {
		fmt.Println("value = ", value)
	}

}

func printArray(array [4]int) {
	// 值拷贝过程
	for _, value := range array {
		fmt.Println("value = ", value)
	}

	array[0] = 111
}
