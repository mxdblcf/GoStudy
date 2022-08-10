package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//使用reader读
func Reader(router gin.Engine) {

	router.GET("/test1", test1)
}
func test1(ctx *gin.Context) {
	response, err := http.Get("http://qiniu.mxdblcf.top/1111.jpg")
	if err != nil || response.StatusCode != http.StatusOK {
		ctx.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	contentType := response.Header.Get("Content-Type")
	contentLength := response.ContentLength
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.jpg"`,
	}
	ctx.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}
