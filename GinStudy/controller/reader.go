package controller

import "github.com/gin-gonic/gin"

//使用reader读
func reader(router gin.Engine) {

	router.GET("/test1", test1)
}
func test1(ctx *gin.Context) {

}
