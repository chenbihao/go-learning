package main

import (
	"fmt"
	"math/rand"
)

// 模拟一次投骰子直到恰好到达 2024 或跳过
func simulate(target int) bool {
	current := 0

	for current < target {
		// 模拟投骰子，随机得到 1 到 6 的点数
		roll := rand.Intn(6) + 1
		current += roll
	}

	// 判断是否恰好到达目标
	return current == target
}

// 打印进度条
func printProgress(current, total int) {
	progress := float64(current) / float64(total)
	barLength := 50                                 // 设置进度条的长度
	numHashes := int(progress * float64(barLength)) // 当前进度的 '#' 数量
	fmt.Printf("\r[")
	for i := 0; i < numHashes; i++ {
		fmt.Print("#")
	}
	for i := numHashes; i < barLength; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("] %.2f%%", progress*100)
}

func main() {

	// 模拟的次数
	totalSimulations := 10000000
	hitTarget := 0

	// 目标点数
	target := 2024

	// 执行模拟
	for i := 0; i < totalSimulations; i++ {
		if simulate(target) {
			hitTarget++
		}
		// 每隔一定的模拟次数更新一次进度
		if i%100 == 0 || i == totalSimulations-1 {
			printProgress(i+1, totalSimulations) // 更新进度条
		}
	}

	// 在进度条结束后，确保换行
	fmt.Printf("\n")

	// 计算恰好到达的概率和跳过的概率
	hitProbability := float64(hitTarget) / float64(totalSimulations)
	missProbability := 1.0 - hitProbability

	// 输出结果
	fmt.Printf("总模拟次数：%d \n", totalSimulations)
	fmt.Printf("恰好到达 %d 的次数：%d  概率：%.10f \n", target, hitTarget, hitProbability)
	fmt.Printf("未恰好到达 %d 的次数：%d  概率：%.10f \n", target, totalSimulations-hitTarget, missProbability)
}
