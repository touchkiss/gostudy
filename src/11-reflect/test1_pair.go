package main

import "fmt"

func main() {
	var a string
	//pair<statictype:string, value:"abcd">
	a = "abcd"
	fmt.Println(a)

	var allType interface{}
	//pair<type:stirng, value:"abcd">
	allType = a
	fmt.Println(allType)

	str, _ := allType.(string)
	fmt.Println(str)
}
