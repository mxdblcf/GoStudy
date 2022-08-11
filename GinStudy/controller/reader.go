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

//权限

func BasicAuthTest(router *gin.Engine) {
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"mxd":    gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}

	//保存几个用户
	authorized := router.Group("/admin", gin.BasicAuth(
		gin.Accounts{
			"mxd":  "dxm",
			"menu": "1234",
		}))

	authorized.GET("/secrets", func(context *gin.Context) {
		//获取由中间件BasicAuth中间价给的用户
		user := context.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": "no secret"})
		}
	})

}
