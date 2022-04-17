package main

import "fmt"

func main() {
	//初始化一个长度为3，容量为5的切片
	numbers := make([]int, 3, 5)
	fmt.Printf("length of numbers = %d ,cap = %d ,numbers = %v\n", len(numbers), cap(numbers), numbers)
	//numbers[20] = 43827
	//fmt.Printf("length of numbers = %d ,cap = %d ,numbers = %v\n", len(numbers), cap(numbers), numbers)
	//向切片追加元素
	numbers = append(numbers, 4)
	fmt.Printf("length of numbers = %d ,cap = %d ,numbers = %v\n", len(numbers), cap(numbers), numbers)
	//向切片追加元素
	numbers = append(numbers, 4)
	fmt.Printf("length of numbers = %d ,cap = %d ,numbers = %v\n", len(numbers), cap(numbers), numbers)
	//向切片追加元素，容量是增加的步长
	numbers = append(numbers, 4)
	fmt.Printf("length of numbers = %d ,cap = %d ,numbers = %v\n", len(numbers), cap(numbers), numbers)
}
