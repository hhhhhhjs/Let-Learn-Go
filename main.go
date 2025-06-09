package main 
// 这里的 main 包名相当于文件夹名，用来引用最外层的包
// 一般来说，在项目中只能有一个 main 包名，用来引用最外层的包
// 这个包是整个程序的入口，main 函数是程序的入口函数，程序从这里开始执行
// 其他的包都是用来被引用的，不能有 main 函数
import (
	"fmt"
	"hhhhhhjs/test"
)


var message string = "go"
func main() {
	fmt.Println("Hello, World!")
	test.Test()
	test.Node()
	// test.FirstFn(message)
}

// init 函数在 package 初始化时执行，并且在 main 函数执行前执行
func init() {
	test.FirstFn(message)
}

// test.FirstFn(message) 
// 这样是错误的执行，会报错，在 go 中函数调用需要发生在任何 函数 的内部，
// 不一定是 main 函数内部，main 函数只是程序的起点

