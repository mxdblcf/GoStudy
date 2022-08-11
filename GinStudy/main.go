package main

import (
	"GoStudy/GinStudy/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @title 测试
// @version 0.0.1
// @description  测试

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//c创建路由组
	group := r.Group("v1")
	{
		group.GET("/m", m)
		group.GET("/n", m2)
	}
	r.GET("/hello", hello)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	//controller.Reader(*r)
	controller.BasicAuthTest(r)
	r.Run()
}

func m(c *gin.Context) {
	//设置一个5秒响应
	//time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{"message": "mm"})
}

func m2(c *gin.Context) {
	c.JSON(http.StatusOK, "1234")
}

// @你好世界
// @Description 你好
// @Accept  json
// @Produce json
// @Success 200 {string} string	"name,helloWorld"
// @Router /hello [get]
func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "helloworldmxd")
}
