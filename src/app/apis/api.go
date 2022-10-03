package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	// "src/app/util"
)

//顶层Api接口，
//当添加新的路由时，继承该Api接口即可
type Api struct {
	Context *gin.Context
}

//OK通常成功时的处理数据
func (a Api) OK(data interface{}, msg string) {
	response.OK(a.Context, data, msg)
}
