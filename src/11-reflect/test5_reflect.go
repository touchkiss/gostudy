package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (this *User) Eat() {
	fmt.Println(this.Name, " is eating...")
}

func reflectFieldAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("input type is ", inputType)

	inputValue := reflect.ValueOf(input)
	fmt.Println("input value is ", inputValue)

	//	通过type获取里面的字段
	//	1.获取interface的reflect.Type，通过Type得到NumField，进行遍历
	//	2.得到每个field，数据类型
	//	3.通过field的Interface()方法得到value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Println("field : ", field.Name, " ,value is ", value)
	}

	//	通过type获取对象的方法
	//	高版本go不支持NumMethod
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Println("method is ", method.Name)
	}
}

func main() {
	xiaoming := User{"小明", 15}
	reflectFieldAndMethod(xiaoming)
}
