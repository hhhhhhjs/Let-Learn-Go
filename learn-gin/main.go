package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
)
// Go 结构体必须大写才能被导出，而 JSON/数据库通常使用小写字段，所以定义 struct 用来解决两者差异问题
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// 定义专辑切片
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


// getAlbums 以 JSON 格式响应所有专辑的列表.
func getAlbums(c *gin.Context) {
	// c 是 gin 提供的 http 上下文对象 context, *gin.Context 是类型

    // Go 的检查比较严格，返回的数据必须是一个合法的 go 数据类型(结构体，切片，map等)
    // c.JSON(http.StatusOK, map[string]string{"name":"小黄"}) // 返回 map
    c.JSON(http.StatusOK, albums) // 返回结构体
	// c.IndentedJSON 是 gin 中提供的一个方法用来将 go 中的对象(结构体、切片、map编码为格式化的 JSON 并作为响应返回)
}

func postAlbums(c *gin.Context) {
    var newAlbum album
    
    // 方法中需要传递变量的指针地址
    // c.BindJSON 方法将前端传递的请求体中的数据反序列化到 newAlbum 中
    if err := c.BindJSON(&newAlbum); err != nil {
        return 
    }

    // 将前端传递的 json 添加到现在的专辑上
    albums = append(albums, newAlbum)

    // statusok 就是 200 这个状态码，只不过这样写更加符合语义化，官方也更推荐这种写法
    c.JSON(http.StatusOK, albums) // 这里直接返回了更新过后数据
}

func getAlbumsById(c *gin.Context) {
    id := c.Param("id")
    
    // 遍历切片
    for index, value := range albums {
        if value.ID == id {
            c.JSON(http.StatusOK, value)
            fmt.Println("index", index)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message":"not found this message"})
    // 这里的 gin.H 是 gin 框架提供的一个小工具，用来直接返回一个 map, 等价于 map[sting]interface{}
    // interface{} 在 Go 中代表空接口,类似于 TS 中的 any
}
func main() {
    router := gin.Default() // 初始化 gin 路由器
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumsById)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}
