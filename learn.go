// 声明包
package main

// 引入包
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

// 学习gin的main函数
func main() {
	fmt.Println("run gin")
	
	// 生成一个实例，为WSGI应用
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 根目录下直接返回字符串
	r.GET("/", func(c *gin.Context){
		c.String(http.StatusOK, "who are you ?")
	})
	// 获取路由参数
	r.GET("/api/:name", func(c *gin.Context){
		name := c.Param("name")
		c.String(http.StatusOK, "you are %s", name)
	})
	// 获取query参数
	r.GET("/api/query", func(c *gin.Context){
		pages := c.Query("pages")
		size := c.Query("size")
		cate := c.DefaultQuery("cate", "MATH")
		c.String(http.StatusOK, "pages %s", pages)
		c.String(http.StatusOK, "size %s", size)
		c.String(http.StatusOK, "cate %s", cate)
	})
	// 获取post参数
	r.POST("/submit", func(c *gin.Context){
		
		//获取参数, 通过PostForm获取的参数值是String类型。
		username := c.PostForm("username")

		// 跟PostForm的区别是可以通过第二个参数设置参数默认值
		pwd := c.DefaultPostForm("pwd", "*****")

		// 区别是GetPostForm返回两个参数，第一个是参数值，第二个参数是参数是否存在的bool值，可以用来判断参数是否存在。
		id, ok := c.GetPostForm("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"massage": "no ok",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"pwd": pwd,
			"id": id,
		})
	})
	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})
	
	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})
	// 分组路由
	defaultHandler := func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	// group: v1
	v1 := r.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	// group: v2
	v2 := r.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}
	r.GET("/testProto", func(c *gin.Context){
		test := &Student{
			Name: "geektutu",
			Male:  true,
			Scores: []int32{98, 85, 88},
		}
		fmt.Printf("\ntest: %v", test)

		// 序列化
		data, err := proto.Marshal(test) 
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}
		fmt.Printf("\nMarshal data: %v", data)

		newTest := &Student{}
		// 反序列化
		err = proto.Unmarshal(data, newTest)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		// Now test and newTest contain the same data.
		if test.GetName() != newTest.GetName() {
			log.Fatalf("data mismatch %q != %q", test.GetName(), newTest.GetName())
		}

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// 函数 func定义函数，+函数名
// func main() {
// 	// 变量定义
// 	// var [name] [type] = ...
// 	var str string = "123"
// 	fmt.Println(str)

// 	// 同时定义两个变量
// 	var a, b int = 1, 2
// 	fmt.Println(a)
// 	fmt.Println(b)

// 	// 先定义变量，在赋值
// 	var c bool
// 	c = true
// 	fmt.Println(c)

// 	// 声明变量语法糖
// 	d := 1
// 	fmt.Println(d)
	
// 	task()
// 	taskAssign()
	
// 	res1 := 1
// 	res2 := 0
// 	swapByVal(res1, res2)
// 	fmt.Println(res1, res2)

// 	swapByRef(&res1, &res2)
// 	fmt.Println(res1, res2)
// 	// 进行不同类型的定义
// 	// var a *int -> 整型指针
// 	// var a []int -> 整型列表
// 	// var a map[string]int -> 字典（key<string>, value<int>）
// 	// var a chan int -> 整型channel
// 	// var a func(string) int -> 定义函数，入参数string，出参数int
// 	fmt.Println()
// 	oneDimArray()
// 	multiDimArray()
// 	newStruct()

// 	john := Student{"john", "muston", 34, "male", true}
// 	fmt.Println(getStuInfo(john))
// 	newSlice()

// }
