package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("----- foo1 -----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c
}

// 多返回值
func foo2(a string, b int) (int, int) {
	fmt.Println("----- foo2 -----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

//返回多个返回值，有形参名称的
//r1,r2属于返回的形参，默认值0
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----- foo3 -----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

//多个返回值形参类型相同，可省略前面的类型
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("----- foo4 -----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//未赋值前已有默认值0
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	r1 = 10000
	r2 = 20000
	return
}

func main() {
	c := foo1("slkdj", 234)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("sdkljfh", 7236)
	fmt.Println("ret1 = ", ret1, ", ret2 = ", ret2)

	ret1, ret2 = foo3("ksjdf", 28376)
	fmt.Println("ret1 = ", ret1, ", ret2 = ", ret2)

	ret1, ret2 = foo4("ksjdf", 28376)
	fmt.Println("ret1 = ", ret1, ", ret2 = ", ret2)

}
