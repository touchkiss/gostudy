package main

import (
	// 前面加下划线表示匿名导入，可执行包里的init函数，但不可调用，否则导入不使用该包会报错
	_ "gostudy/src/5-init/lib1"
	"gostudy/src/5-init/lib2"
	//前面加点表示导入该包里的所有方法，可以直接执行方法名
	//. "gostudy/src/5-init/lib2"
	// 表示给导入的包起别名aa，可用aa调用包里的方法
	aa "fmt"
)

func main() {
	//lib1.Lib1Test()
	lib2.Lib2Test()
	//Lib2Test()
	aa.Println("这里是别名调用")
}
