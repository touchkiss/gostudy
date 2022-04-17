package main

import "fmt"

//如果类名大写表示其他包也能够访问
type Hero struct {
	//属性首字母大写表示可以对外访问，否则只能在类的内部访问
	Name  string
	Age   int
	level int
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(name string) {
	this.Name = name
}

func main() {
	hero1 := Hero{"小明", 12, 4}

	fmt.Println("name = ", hero1.GetName())

	fmt.Println("===========")
	hero1.SetName("小红")
	fmt.Println("name = ", hero1.GetName())
}
