package routers

import (
	"github.com/gin-gonic/gin"
	"src/app/service"
)

func init() {
	RouterList = append(RouterList, UserController)
}

func UserController(g *gin.RouterGroup) {
	group := g.Group("/user")
	userService := service.User{}
	{
		//group.GET("/info", userService.GetInfo)
		group.GET("/edit", userService.EditName)
		group.GET("/info", userService.GetInfo)
		group.POST("/post", userService.PostInfo)
		group.GET("/")
	}
}
