package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "src/docs"
)

type RouterInit interface {
	Init()
	CheckAllRoute(r *gin.Engine)
}

type Router struct {
}

var (
	RouterList = make([]func(group *gin.RouterGroup), 0)
)

func (r Router) Init() {
	e := gin.Default()
	r.CheckAllRoute(e)
	//handle swagger
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	e.Run()
}

func (r Router) CheckAllRoute(e *gin.Engine) {
	//根据需要更改顶层路由
	group := e.Group("/api")
	for _, f := range RouterList {
		f(group)
	}
}
