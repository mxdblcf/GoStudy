package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {

	// 创建一个默认的路由引擎
	r := gin.Default()

	r.GET("/m", m)
	//	r.PUT("d",p)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	//controller.Reader(*r)
	//	controller.BasicAuthTest(r)
	r.Run()

}

//func p(context *gin.Context) {
//
//	body := context.Request.Body
//	f, e := os.Create(os.Getenv("oss") + "/objects/" +
//		strings.Split(r.URL.EscapedPath(), "/")[2])
//	if e != nil {
//		log.Println(e)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	defer f.Close()
//	io.Copy(f, context.)
//}

func m(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"a": os.Getenv("PATH")})
}
