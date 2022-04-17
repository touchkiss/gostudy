package main

import "fmt"

//声明全局变量 方法一、二、三是可以的
var gA int = 100
var gB = 200

func main() {
	//	方法一：声明一个变量，默认值是0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of a = %T\n", a)

	//	方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s,type of bb = %T\n", bb, bb)

	//	方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s,type of cc = %T\n", cc, cc)

	//	方法四：（常用的方式）省去var关键字，直接自动匹配
	//	只能在函数体内使用，不可用于全局变量
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	ee := "abcd"
	fmt.Printf("ee = %s,type of ee = %T\n", ee, ee)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("g = %f,type of g = %T\n", g, g)

	//	声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, " , yy = ", yy)

	var kk, ll = 100, "abfsjl"
	fmt.Println("kk = ", kk, " , ll = ", ll)

	//	多行的多变量声明
	var (
		vv int  = 200
		jj bool = true
	)
	fmt.Println("vv = ", vv, " , jj = ", jj)
}
