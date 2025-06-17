# 注意点
## 规则
* 在 Go 中，Go 不允许在同一个文件夹中定义多个包名，所有文件必须使用**相同的包名**
* Go 的导入路径要求被导入的包必须在一个**单独的文件夹**中
* Go 的规则比较严格，如果一个包引入了 fmt 但是没有使用，会报错
* 同一个包中的函数可以直接相互调用

## 编译并安装应用程序
* 首先，使用 go build 将代码编译成可执行文件
* 接下来需要安装可执行文件，以便运行他，目前也可以运行，但是如果想要运行的话，只能在当前目录运行；
* 执行 go list -f '{{.Target}}', 这会输出 Go 的安装路径,
`expect output: /home/gopher/bin/hello`, 这证明 Go 的安装目录是 /home/gopher/bin
* ```$set PATH=%PATH%;C:\path\to\your\install\directory 使用上述命令将二进制安装到/home/gopher/bin 现现在自己的系统安装命令示例:set PATH=%PATH%;C:\Users\15124\go\bin ```
* ```设置好安装目录之后，使用 go install 命令来编译和安装包，现在即可在任何地方运行这个包文件 ```