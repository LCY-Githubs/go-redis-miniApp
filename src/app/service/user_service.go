package service

import (
	"github.com/gin-gonic/gin"
	"src/app/util"
)

type User struct {
}

type UserService interface {
	GetInfo(*gin.Context)
	EditName(*gin.Context)
	PostInfo(ctx *gin.Context)
}

// GetInfo GetPage
// @Summary 用户信息
// @Description 得到信息
// @Tags 编辑名称
// @Param name query string false "name"
// @Router /api/user/info [get]
// @Security Bearer
func (u User) GetInfo(c *gin.Context) {
	util.OK(c, "success", nil)
}

/*
/api/v0/dept [get]
*/

// EditName GetPage
// @Summary 用户信息
// @Description 分页列表
// @Tags 编辑名称
// @Param name query string false "name"
// @Router /api/user/edit [get]
// @Security Bearer
func (u User) EditName(c *gin.Context) {
	param := c.Query("name")
	println(param)
	util.OK(c, "getName", "yes")
}

// PostInfo GetPage
// @Summary 用户信息提交
// @Description postTest
// @Tags 测试post
// @Param message formData string false "name"
// @Router /api/user/post [post]
// @Produce json
// @success 200 {string} string
func (u User) PostInfo(c *gin.Context) {
	form := c.PostForm("message")
	println(form)
	util.OK(c, "post Success", nil)
}
