package lib1

import "fmt"

//当前lib1包提供的api
//对外开放的方法首字母大写，私有方法首字母小写
func Lib1Test() {
	fmt.Println("Lib1Test()")
}

func init() {
	fmt.Println("lib1 init")
}
