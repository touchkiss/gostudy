package main

import "fmt"

func main() {
	//	声明slice1是一个切片，并且初始化

	slice1 := []int{1, 2, 3}

	fmt.Printf("len = %d, slice1 = %v\n", len(slice1), slice1)

	//声明slice2是一个切片，但是并没有分配空间
	var slice2 []int
	//开辟3个空间
	slice2 = make([]int, 3)
	slice2[0] = 100
	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)
	//slice2[20] = 299
	//fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)

	var slice3 []int
	//判断slice3是一个空切片
	fmt.Println(slice3 == nil)

	//声明slice4是一个切片，同时给slice4分配空间，默认初始化值为0
	slice4 := make([]int, 5)
	fmt.Printf("len = %d, slice4 = %v\n", len(slice4), slice4)
	slice4[2] = 30
	fmt.Printf("len = %d, slice4 = %v\n", len(slice4), slice4)
}
