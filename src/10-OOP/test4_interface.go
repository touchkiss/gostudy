package main

import "fmt"

//本质是一个指针
type Animal interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的类型
}

//通过实现interface的所有方法，达成实现interface的目的，少写则不算
type Cat struct {
	color string //颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "cat"
}

//通过实现interface的所有方法，达成实现interface的目的，少写则不算
type Dog struct {
	color string //颜色
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleeping")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "dog"
}

func Walking(animal Animal) {
	fmt.Println("type is ", animal.GetType(), " is walking")
}

func main() {
	mimi := Cat{"yellow"}
	mimi.Sleep()
	fmt.Println("color is ", mimi.GetColor())
	fmt.Println("type is ", mimi.GetType())

	var animal Animal
	animal = &mimi
	animal.Sleep()
	Walking(animal)

	animal = &Dog{"red"}
	animal.Sleep()
	fmt.Println("type is ", animal.GetColor())
	Walking(animal)
}
