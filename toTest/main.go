package main

import (
	"fmt"
)

// 试试 iota
// iota 是一个在 const 常量声明中使用的 预定义标识符，用于生成一组连续的数字。它在每个 const 声明块中 从 0 开始，每新增一行常量声明，iota 的值自动加 1。
// iota 只能在 const 中使用
const (
	Num = iota
	Num1
	Num2
)

// 省略的表达式会自动复用上一行的表达式模板，只是 iota 会加 1
const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)



func main() {
	fmt.Println(Num, Num1, Num2)
	fmt.Println(KB, MB, GB, TB)
}