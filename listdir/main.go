package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建或覆盖名为 "example.txt" 的文件
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // 在函数退出时关闭文件

	// 写入数据到文件
	_, err = file.WriteString("Hello, World!\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created successfully!")
	Run()
}
