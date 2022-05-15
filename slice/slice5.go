package main

import "fmt"

func main()  {
	s := []int{1, 2, 3}

	// [0, 2]
	s1 := s[0:2]  //从 下标0 开始 截取长度为2的元素

	fmt.Println(s1)

	s1[0] = 100

	fmt.Println(s)
	fmt.Println(s1)

	// copy()方法
	s2 := make([]int, 3)

	// 将s中的值依次拷贝到s2中
	copy(s2, s)

	fmt.Println(s2)
}
