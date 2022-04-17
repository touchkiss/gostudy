package main

import (
	"fmt"
	"os"
)

func main() {
	// tty:pair<type:*os.File, value:"C:\Program Files\XMind\XMind.exe"文件描述符>
	tty, err := os.OpenFile("C:\\Users\\MOX\\OneDrive\\桌面\\mcn.har", os.O_RDWR, 0)
	if err != nil {
		fmt.Println(tty)
	}
}
