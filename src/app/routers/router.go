package router

import "github.com/gin-gonic/gin"

//全局变量用于存储所有的路径
var (
	routerSlice = make([]func(*gin.RouterGroup), 0)
)

//加载所有的路由
func InitRouterFunc() {
	//调用所有的路由组
	e := gin.Default()
	CheckAllRoutes(e)
	e.Run()
}

//调用所有的方法
func CheckAllRoutes(r *gin.Engine) {
	//根据需求更改顶层路由
	v1 := r.Group("/api")
	for _, f := range routerSlice {
		f(v1)
	}
}
