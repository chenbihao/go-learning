package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 指定要遍历的目录
	rootDir := `Q:\其他\文件\`
	// 调用遍历函数
	walkDirectorys(rootDir, 1, "", "")

	//walkDirectorysStartJk(rootDir)
}

// 遍历并输出目录
func walkDirectorys(path string, depth int, prefix, suffix string) {
	// 递归结束条件：达到指定的深度
	if depth < 0 {
		return
	}

	// 打印当前目录
	//fmt.Println(path)

	// 读取当前目录下的文件和子目录
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 遍历当前目录下的文件和子目录
	for _, file := range files {
		fmt.Printf("%s%s%s\n", prefix, file.Name(), suffix)
		// 打印子目录下的文件和子目录
		if file.IsDir() {
			walkDirectorys(filepath.Join(path, file.Name()), depth-1, prefix, suffix)
		}
	}
	return
}

// 遍历并输出目录 ,用来把目录转成ob的md类型文本，方便做笔记用
func walkDirectorysStartJk(path string) {

	// 递归结束条件：达到指定的深度
	depth := 1

	// 读取当前目录下的文件和子目录
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 遍历当前目录下的文件和子目录
	for _, file := range files {

		// 截取去掉日期
		var title string
		if index := strings.Index(file.Name(), "第"); index != -1 {
			title = file.Name()[index:]
		} else {
			title = file.Name()
		}
		fmt.Printf("%s%s\n", "### ", title)

		// 打印子目录下的文件和子目录
		if file.IsDir() {
			walkDirectorys(filepath.Join(path, file.Name()), depth-1, "[[", "]]")
		}
	}
	return
}
