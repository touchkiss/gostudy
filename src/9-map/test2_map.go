package main

import "fmt"

func printMap(myMap map[string]string) {
	//引用传递
	//遍历
	for key, value := range myMap {
		fmt.Printf("country : %s capital : %s\n", key, value)
	}
}

func main() {
	capital := make(map[string]string)
	//添加
	capital["CHINA"] = "BEIJING"
	capital["USA"] = "NEW YORK"
	capital["JAPAN"] = "TOKYO"

	printMap(capital)
	fmt.Println("=============")
	//	删除
	delete(capital, "USA")
	printMap(capital)

	fmt.Println("=============")
	//	修改
	capital["JAPAN"] = "GUANDAO"
	printMap(capital)
}
