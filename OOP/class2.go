package main

import "fmt"

// 父类
type Me struct {
	name string
	sex  string
}

func (this *Me) Eat(){
	fmt.Println("Me eat...")
}

func (this *Me) Walk() {
	fmt.Println("Me walk...")
}


// 基于父类方法的重写
func (this *SuperMe) Eat() {
	fmt.Println("SuperMe eat...")
}

func (this *SuperMe) Cool() {
	fmt.Println("SuperMe cool...")
}

//子类
type SuperMe struct {
	Me
	level int
}

func main() {
	me := Me{
		name: "zhangsan",
		sex: "man",
	}

	me.Eat()
	me.Walk()

	//定义一个子类对象

	//superMe := SuperMe{
	//	Me{"lisi", "man"},
	//	level: 88,
	//}

	var superMe SuperMe

	superMe.name = "lisi"
	superMe.sex = "man"
	superMe.level = 100

	// 子类方法调用
	superMe.Eat()
	superMe.Walk()
	superMe.Cool()
}
