package main

import (
	_ "GoStudy/GinStudy/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

// @title 测试
// @version 0.0.1
// @description  测试

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	r.GET("/hello",hello)
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}



// @你好世界
// @Description 你好
// @Accept  json
// @Produce json
// @Success 200 {string} string	"name,helloWorld"
// @Router /hello [get]
 func hello(c *gin.Context) {
		c.JSON(http.StatusOK, "helloworldmxd")}


