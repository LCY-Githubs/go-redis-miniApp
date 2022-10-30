package routers

import (
	"github.com/gin-gonic/gin"
	"src/app/routers"
)

/**
note应该支持便签的

1.新增
2.修改
3.标记为完成
4.按照不同的查询条件进行查看搜索
*/

// @mean 是是是
func init() {
	routers.RouterList = append(routers.RouterList, NoteController)
}

func NoteController(g *gin.RouterGroup) {
	g.Group("/note")
	{
		g.POST("/add", nil)
		g.POST("edit", nil)
		g.POST("/markFinish", nil)
		g.POST("/get", nil)
	}
}
