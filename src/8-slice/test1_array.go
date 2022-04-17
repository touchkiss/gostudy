package main

import "fmt"

func main() {
	//	固定长度的数组，默认0
	var myArray1 [10]int
	for i := 0; i < 10; i++ {
		fmt.Println(myArray1[i])
	}

	//	初始化部分值
	myArray2 := [10]int{1, 2, 3, 4, 5}
	for i := 0; i < len(myArray2); i++ {
		fmt.Println(myArray2[i])
	}
	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	myArray3 := [4]int{1, 2, 3, 4}
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)
	fmt.Println(myArray3)
}
