package testError1 
// 这里的 package name 一般需要跟文件夹名一致，这样有利于后期开发维护
// 也可以不维护，但是在 main 函数中调用时，需要使用包名，所以文件夹名字和包名一致是最佳实践

import (
	"fmt"
	"errors"
)

func GetName(name string) (string, error) {
	if(name == ""){
		return "", errors.New("empty name")
	}

	var message string
	// 这里 %v 为占位符，代表后面的变量
	// Sprintf 为直接返回字符串但是不打印
	message = fmt.Sprintf("Hi, %v. welcome", name)
	return message, nil // nil 是 go 中代表返回的 空值
}