package main
import (
	"fmt"
	"greetings/arraySlice"
)

// 切片
// 切片的长度是动态的，可以使用 append 方法进行添加元素
func main() {
	var slice []string
	fmt.Println(slice) // expeted: []

	fmt.Println("arrayslice")
	arraySlice.Slice()
	arraySlice.Array()
}

