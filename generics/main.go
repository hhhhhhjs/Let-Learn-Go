package main

import (
	"fmt"
)

// 使用接口
type Number interface {
	int64 | float64
}
// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m { // 注意: 使用 range 遍历 map 时是无序遍历
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	// 其中 K 是 map 的 key 的类型，必须是 comparable, 可比较的
	// V 是 map 的 value 类型, 是 int64 或者 float64 类型
    var s V
    for _, v := range m {
        s += v
    }
    return s
}


func main() {
	// 1. 不使用泛型，需要声明两个函数
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))
	// 2. 使用泛型

	// 函数调用时可以不加泛型参数，编译器可以自行推断
	fmt.Printf("Generic Sums: %v and %v\n",
    SumIntsOrFloats[string, int64](ints),
    SumIntsOrFloats[string, float64](floats))

	// 删除泛型参数
	fmt.Printf("Generic Sums: %v and %v\n",
    SumIntsOrFloats(ints),
    SumIntsOrFloats(floats))

}