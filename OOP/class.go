package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

// 定义class

// (this.Hero) 表示当前的方法绑定在结构体Hero上

// 这是一个hear的值拷贝
/*func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this Hero) getName() string {
	return this.Name
}

func (this Hero) setName() {
	this.Name = "lisi"
}*/

// 使用指针才能set hear对象里面的数据
func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this *Hero) getName() string {
	return this.Name
}

func (this *Hero) setName(newName string) {
	this.Name = newName
}

func main()  {
	hear := Hero{
		Name: "zhangsan",
		Ad: 100,
		Level: 1,
	}

	hear.Show()

	hear.getName()

	hear.setName("lisi")

	fmt.Println("===========")

	hear.Show()
}
