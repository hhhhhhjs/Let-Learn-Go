package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"html/template"
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

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	p, _ := loadPage(title)
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}

// 修改 viewHandle 使用 html/template
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	// 在 go 中,不需要手动 import 引入文件，只需要引入 package，go 会自动找到文件
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}



// func editHandler(w http.ResponseWriter, r *http.Request) {
//     title := r.URL.Path[len("/edit/"):]
//     p, err := loadPage(title)
//     if err != nil {
//         p = &Page{Title: title}
//     }
//     fmt.Fprintf(w, "<h1>Editing %s</h1>"+
//         "<form action=\"/save/%s\" method=\"POST\">"+
//         "<textarea name=\"body\">%s</textarea><br>"+
//         "<input type=\"submit\" value=\"Save\">"+
//         "</form>",
//         p.Title, p.Title, p.Body)
// }

// 使用 html/template
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
	// 在 go 中,不需要手动 import 引入文件，只需要引入 package，go 会自动找到文件
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
}



func main() {
	// 服务器启动之后会阻塞程序，需要放在服务器启动之前执行
	// var result = test.Test()
	// fmt.Println(result)
	http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
	// 这里相当于 new 一个实例
	// fmt.Println("Hello, World!")
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// // 使用实例对象上的方法
	// // 写文件，创建一个名为 TestPage 的 txt 文本文件从、
	// p1.save()
	// // 读文件，读取文件中的内容
	// p2, _ := loadPage("TestPage")
	// // 目的是想要打印一下文件中的内容
	// fmt.Println(string(p2.Body))
	log.Fatal(http.ListenAndServe(":8080", nil))
}