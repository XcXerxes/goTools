package tools

import (
	"fmt"
	"os"
)

func WriteAndRead()  {
	// 写入
	file, err := os.OpenFile("../test.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0x644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	wint, err := file.WriteString("hello world")
	if err != nil {
		panic(err)
	}
	fmt.Println("content============", wint)

	// 读取
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	bs := make([]byte, 100)
	rint, err := file.Read(bs)
	if err != nil {
		panic(err)
	}
	fmt.Println("=====================", rint, bs)
}
