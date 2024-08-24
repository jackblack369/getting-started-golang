package main

import (
	"fmt"
)

// 🤔定义一个类型为函数的别名
type IntTransformer func(int) int

// 将多个转换函数组合成一个管道
func pipe(value int, transformers ...IntTransformer) int {
	for _, transformer := range transformers {
		value = transformer(value)
	}
	return value
}

// 定义一些转换函数
func addOne(x int) int {
	return x + 1
}

func square(x int) int {
	return x * x
}

func main() {
	// 使用管道处理数据
	result := pipe(3, addOne, square)
	fmt.Println("Result:", result) // 输出 Result: 16
}
