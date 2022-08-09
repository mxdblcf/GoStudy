package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	hertz := server.Default()
	hertz.GET("/ping",
		func(c context.Context, ctx *app.RequestContext) {
			ctx.JSON(consts.StatusOK, utils.H{"ping": "png"})
		})
	hertz.Spin()
}
