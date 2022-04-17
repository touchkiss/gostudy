package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	//	[0,1]
	s1 := s[0:1]
	fmt.Printf("length of s1 = %d ,s1 = %v\n", len(s1), s1)

	s2 := make([]int, 2)
	//copy函数，将s复制给s2
	copy(s2, s)
	fmt.Printf("length of s2 = %d ,s2 = %v\n", len(s2), s2)
}
