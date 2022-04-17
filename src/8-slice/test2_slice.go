package main

import "fmt"

// slice是引用传递，而array是值传递
func printArray(myArray []int) {
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4} //动态数组，slice切片
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
	printArray(myArray)

	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
}
