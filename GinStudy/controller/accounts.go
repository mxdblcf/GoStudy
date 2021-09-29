package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShowAccount godoc
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Success 200  {string} Helloworld4
// @Failure 400 {string} Helloworld3
// @Failure 404 {string} Helloworld2
// @Failure 500 {string} Helloworld1
// @Router /accounts/mxd [get]
func (c *Controller) ShowAccount(ctx *gin.Context)  {
ctx.JSON(http.StatusOK,"mxd")
}
