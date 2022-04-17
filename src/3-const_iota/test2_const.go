package main

import "fmt"

//const 定义枚举类型
const (
	//	可以在const（）添加一个关键字iota，每行的iota都会累加1，第一行的iota默认值是0
	//  iota只能够配合const（）使用，不能用于其他方法
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2 //iota=0, a = iota + 1 = 1, b = iota + 2 = 2
	c, d                      // iota = 1, c = iota + 1 = 2, d = iota + 2 = 3
	e, f                      // iota = 2, e = iota + 1 = 3, f = iota + 2 = 4

	g, h = iota * 2, iota * 3 // iota = 3, g = iota * 2 = 6, h = iota * 3 = 9
	l, m                      // iota = 4, l = iota * 2 = 8, m = iota * 3 = 12
)

func main() {
	//	常量（只读属性）
	const length int = 10

	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a, " ,b = ", b)
	fmt.Println("c = ", c, " ,d = ", d)
	fmt.Println("e = ", e, " ,f = ", f)

	fmt.Println("l = ", l, " ,m = ", m)
}
