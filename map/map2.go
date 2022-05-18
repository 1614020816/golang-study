package main

import "fmt"

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["sichuan"] = "chengdu"
	cityMap["guangdong"] = "shenzhen"
	cityMap["gansu"] = "lanzhou"

	// 遍历
	printMap(cityMap)

	// 修改
	cityMap["guangdong"] = "guangzhou"

	// 删除
	delete(cityMap, "gansu")

	fmt.Println("=============")

	changeMap(cityMap)

	// 遍历
	printMap(cityMap)
}

func printMap(cityMap map[string]string){
	// cityMap 是一个引用类型， 实际上是传递的cityMap的指针
	for key, value := range cityMap{
		fmt.Println("key = ", key, ", value = ", value)
	}
}

func changeMap(cityMap map[string]string){
	cityMap["shanxi"] = "xian"
}