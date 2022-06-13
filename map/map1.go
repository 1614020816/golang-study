package main

import "fmt"

func main() {
	// 第一种申明方式
	var myMap1 map[string] string

	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	if len(myMap1) == 0 {
		fmt.Println("myMap1 是一个空map")
	}

	// 在使用时 要先给map开辟数据空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "js"
	myMap1["two"] = "html"
	myMap1["three"] = "css"

	fmt.Println(myMap1)

	//  第二种声明方式
	myMap2 := make(map[int]string)

	myMap2[1] = "js"
	myMap2[2] = "html"
	myMap2[3] = "css"

	fmt.Println(myMap2)

	myMap3 := map[int]string{
		1: "js",
		2: "html",
		3: "css",
	}

	fmt.Println(myMap3)
}
