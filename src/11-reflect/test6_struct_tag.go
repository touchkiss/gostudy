package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	//将这种方式 理解为java中的注解
	Name string `info:"name" doc:"我的名字"`
	Age  int    `info:"age"`
}

func findTag(arg interface{}) {
	t := reflect.TypeOf(arg).Elem()
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info : ", taginfo, " doc : ", tagdoc)
	}
}

func main() {
	rs := resume{"小黄", 15}

	findTag(&rs)
}
