package tools

import (
	"fmt"
	"strconv"
	"strings"
)

// ParsePageRange 解析页面范围字符串并返回页码的切片
func ParsePageRange(pageRange string) ([]int, error) {
	if pageRange == "" {
		return nil, fmt.Errorf("空参数")
	}
	var pages []int
	for _, part := range strings.Split(pageRange, ",") {
		if strings.Contains(part, "-") {
			parts := strings.Split(part, "-")
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, fmt.Errorf("无效起始页: %s", parts[0])
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("无效结束页: %s", parts[1])
			}
			for i := start; i <= end; i++ {
				pages = append(pages, i)
			}
		} else {
			page, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("无效页面: %s", part)
			}
			pages = append(pages, page)
		}
	}
	return pages, nil
}
