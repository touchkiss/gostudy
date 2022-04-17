package main

import "fmt"

func func1() {
	fmt.Println("func1()")
}

func func2() {
	fmt.Println("func2()")
}

func func3() {
	fmt.Println("func3()")
}

func returnFunc() string {
	fmt.Println("returnFunc()")
	return "returnFunc"
}

//defer 会在return之后执行
func func4() string {
	defer fmt.Println("defer func")
	return returnFunc()
}

func main() {
	//defer 会在函数的最后执行，压栈规则，先写的后执行
	//下为输出结果
	//main()
	//func3()
	//func2()
	//func1()

	defer func1()
	defer func2()
	defer func3()

	fmt.Println("main()")

	fmt.Println(func4())
}
