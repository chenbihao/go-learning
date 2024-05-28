package tools

// InSliceStr 判断字符串是否在 slice 中。
func InSliceStr(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// InSliceInt 判断int是否在 slice 中。
func InSliceInt(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// GetSliceMaxInt 获取切片最大max
func GetSliceMaxInt(items []int) (maxVal int) {
	for _, val := range items {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
