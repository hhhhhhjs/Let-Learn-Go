package ArraySlice

import (
	"fmt"
)

// 切片
func Slice() {
	var slice []string
	slice = append(slice, "哈哈")
	slice = append(slice, "嘿嘿")
	fmt.Println(slice)
}

func Array() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)
	// 数组的长度是固定的，不能改变
	// 数组的长度可以使用 len 函数获取
	fmt.Println(len(arr))
	// 数组的元素可以使用索引来访问
	fmt.Println(arr[0])
	// 数组的元素可以使用 for 循环来遍历
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
