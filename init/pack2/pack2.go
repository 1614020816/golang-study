package pack2

import "fmt"
import "../pack1"

func init() {
	fmt.Println("Pack2 init()......")
}

// 当前包的对外api，函数名首字母大写表示对外开放的函数，小写表示只能是在本文件内调用
func Pack2Test() {
	pack1.Pack1Test()
	fmt.Println("Pack2 Pack2Test()......")
}