package main

import "fmt"

type Life struct {
	Type string
}

type Man struct {
	Name string
	Age  int
}

func (this *Man) Walk() {
	fmt.Println(this.Name, " walking。。。")
}

func (this *Man) eat() {
	fmt.Println(this.Name, " eating...")
}

type SuperMan struct {
	Man   //SuperMan继承了Man
	Life  //多继承
	level int
}

//重定义父类方法
func (this *SuperMan) eat() {
	fmt.Println("this is superMan ", this.Name, " ,level is ", this.level)
}

func (this *SuperMan) power() {
	fmt.Println("this type is ", this.Type)
}

func main() {
	man1 := Man{"小红", 15}
	man1.Walk()
	man1.eat()

	superMan1 := SuperMan{Man{"超人", 12}, Life{"dlfks"}, 5}
	superMan1.eat()
	superMan1.Walk()
	superMan1.power()
}
