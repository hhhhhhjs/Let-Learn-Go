package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

// 这里相当于 js 中的 class Page { save() {}}, save 就是类中内置的一个方法
func (p *Page) save() error {
	// p *Page 是方法的接收者
	// 这里 需要使用指针 *Page ，就像 js 中引用数据类型引用的是对象的地址一样
	// 如果使用 Page 则相当于拷贝了一份 Page, 修改之后并不会影响原对象
	filename := p.Title + ".txt"
	// 这里没有给文件路径，默认在文件的运行目录底下
	return os.WriteFile(filename, p.Body, 0600) // 如果写入成功返回 nil,如果失败则返回错误信息
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 这里 w 作为第一个参数的作用是为了告诉编译器要通过 w 来将内容发送给浏览器
	// w 是一个管道，stream
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func main() {
	http.HandleFunc("/", handler)
	// 这里相当于 new 一个实例
	fmt.Println("Hello, World!")
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// 使用实例对象上的方法
	// 写文件，创建一个名为 TestPage 的 txt 文本文件从、
	p1.save()
	// 读文件，读取文件中的内容
	p2, _ := loadPage("TestPage")
	// 目的是想要打印一下文件中的内容
	fmt.Println(string(p2.Body))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
