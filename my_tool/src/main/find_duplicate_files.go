package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// 查找重复文件
func main() {
	rootDir := `Q:\其他\文件\`
	findDuplicateFiles(rootDir)
}

// 查找重复文件
func findDuplicateFiles(rootDir string) {
	fileSizeMap := make(map[int64][]string)

	// 遍历文件夹
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// 记录文件大小和路径
			size := info.Size()
			fileSizeMap[size] = append(fileSizeMap[size], path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("遍历文件夹时出错:", err)
		return
	}

	// 提取并排序文件大小
	sizes := make([]int64, 0, len(fileSizeMap))
	for size := range fileSizeMap {
		sizes = append(sizes, size)
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	// 输出重复的文件（按文件大小从大到小）
	for _, size := range sizes {
		paths := fileSizeMap[size]
		if len(paths) > 1 {
			fmt.Printf("-------------------------\n")
			fmt.Printf("文件大小 %s KB 的重复文件:\n", formatKB(size))
			for _, path := range paths {
				fmt.Println("	" + path)
			}
		}
	}
}

func formatKB(bytes int64) string {
	kb := (bytes + 1023) / 1024 // 使用整数除法并四舍五入
	return fmt.Sprintf("%d KB", kb)
}
