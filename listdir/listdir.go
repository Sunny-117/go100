package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListDir(dirPath string, indent int) {
	// 获取目录下的所有文件和子目录
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", dirPath, err)
		return
	}

	// 打印当前目录的内容
	for _, file := range files {
		// 添加缩进
		for i := 0; i < indent; i++ {
			fmt.Print("  ")
		}

		// 打印文件或目录名
		fmt.Println(file.Name())

		// 如果是目录，则递归打印其内容
		if file.IsDir() {
			subdirPath := filepath.Join(dirPath, file.Name())
			ListDir(subdirPath, indent+1)
		}
	}
}

func Run() {
	// 检查命令行参数，确保提供了目录路径
	if len(os.Args) != 2 {
		fmt.Println("Usage: listdir <directory>")
		os.Exit(1)
	}

	// 获取目录路径
	dirPath := os.Args[1]

	// 调用函数开始打印目录列表
	ListDir(dirPath, 0)
}
