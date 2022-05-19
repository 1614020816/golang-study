package main

import "fmt"

// 定义一个接口，本质是指针
type IAnimal interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func ShowAnimal(animal IAnimal){
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {
	var animal IAnimal // 接口数据类型，父类指针

	animal = &Cat{"Orange"}

	animal.Sleep()

	animal = &Dog{"Yellow"}

	animal.Sleep()

	cat := &Cat{"Orange"}
	dog := &Dog{"Yellow"}

	ShowAnimal(cat)
	fmt.Println("==========")
	ShowAnimal(dog)

}



