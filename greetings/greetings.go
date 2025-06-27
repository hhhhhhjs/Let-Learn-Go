package main
import (
	"fmt"
	"greetings/arrayslice"
)

// 切片
// 切片的长度是动态的，可以使用 append 方法进行添加元素
func main() {
	var slice []string
	fmt.Println(slice) // expeted: []

	fmt.Println("arrayslice")
	arrayslice.Slice()
	arrayslice.Array()
}

