package main

import "fmt"

func main() {
	// 1. 声明slice是一个切片，并且初始化，默认值是1， 2， 3. 长度len是3
	//slice := []int{1, 2, 3}

	// 2. 声明slice是一个切片，但是并没有给slice分配空间
	//var slice []int
	//slice = make([]int, 3)  // 开辟3个空间，默认值为0

	// 3. 声明slice是一个切片, 同时给slice分配空间，3个空间，初始值是0
	//var slice []int = make([]int ,3)

	// 4. 声明slice是一个切片, 同时给slice分配空间，3个空间，初始值是0，通过 := 推导出slice是一个切片
	slice := make([]int, 3)

	fmt.Printf("len = %d, slice = %v\n", len(slice), slice)

	// 判断一个slice是否为0
	if slice == nil {
		fmt.Println("slice 是一个空切片")
	} else {
		fmt.Println("slice 是有空间的")
	}
}
