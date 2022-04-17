package main

import "fmt"

type myInt int

//定义一个结构体
type Book struct {
	title string
	auth  string
	price int
}

//值传递，不改变
func changeBook(book Book) {
	book.title = "值传递，不改变"
}

func changeBook2(book *Book) {
	book.title = "引用传递，改变"
}

func main() {
	var book1 Book
	book1.title = "哈哈哈"
	book1.auth = "小王"
	book1.price = 32
	fmt.Printf("%v\n", book1)
	changeBook(book1)
	fmt.Printf("%v\n", book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
