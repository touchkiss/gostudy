package main

import "fmt"

//interfact{}是万能类型，类似于java的Object
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called")
	fmt.Println(arg)

	//interface{} 该复合区分此时的底层数据类型是什么
	//给interface{}提供“类型断言”机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string")
		fmt.Println("value is ", value)
	}
}

func main() {
	myFunc(23)

	myFunc("abc")
}
