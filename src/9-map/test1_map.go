package main

import "fmt"

func main() {
	//声明myMap1是一种map类型，key是string类型，value是int
	var myMap1 map[string]int
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}
	//使用map前，用make分配空间
	myMap1 = make(map[string]int)
	myMap1["one"] = 10
	myMap1["two"] = 987
	fmt.Printf("length of myMap1 = %d, myMap1 = %v\n", len(myMap1), myMap1)

	fmt.Println("one = ", myMap1["one"])

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "python"
	fmt.Printf("length of myMap2 = %d, myMap2 = %v\n", len(myMap2), myMap2)
	fmt.Println("1 = ", myMap2[1])

}
